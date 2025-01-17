package rpc_test

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	mathrand "math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/go-sequence/intents"
	"github.com/0xsequence/nitrocontrol/enclave"
	"github.com/0xsequence/waas-authenticator/data"
	"github.com/0xsequence/waas-authenticator/proto"
	"github.com/0xsequence/waas-authenticator/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRPC_SendIntent_SendTransaction(t *testing.T) {

	block, _ := pem.Decode([]byte(testPrivateKey))
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.NoError(t, err)

	cfg := initConfig(t)

	issuer, _, closeJWKS := issueAccessTokenAndRunJwksServer(t)
	defer closeJWKS()

	random := mathrand.New(mathrand.NewSource(42))
	kmsClient := &kmsMock{random: random}
	enc, err := enclave.New(context.Background(), enclave.DummyProvider, kmsClient, privKey)
	require.NoError(t, err)

	sessWallet, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)
	signingSession := intents.NewSessionP256K1(sessWallet)

	tenant, tntData := newTenant(t, enc, issuer)
	acc := newAccount(t, enc, issuer, sessWallet)
	sess := newSession(t, enc, issuer, signingSession)

	walletAddr, err := rpc.AddressForUser(context.Background(), tntData, acc.UserID)
	require.NoError(t, err)

	dbClient := &dbMock{
		sessions: map[string]*data.Session{sess.ID: sess},
		tenants:  map[uint64][]*data.Tenant{tenant.ProjectID: {tenant}},
	}
	svc := initRPC(cfg, enc, dbClient)

	srv := httptest.NewServer(svc.Handler())
	defer srv.Close()

	intentData := &intents.IntentDataSendTransaction{
		Identifier: "identifier",
		Wallet:     walletAddr,
		Network:    "1",
		Transactions: []json.RawMessage{
			json.RawMessage(`{"data":"0x010203","to":"0x27CabC9700EE6Db2797b6AC1e1eCe81C72A2cD8D","type":"transaction","value":"0x2000000000"}`),
		},
	}
	intent := generateSignedIntent(t, intents.IntentName_sendTransaction.String(), intentData, signingSession)

	c := proto.NewWaasAuthenticatorClient(srv.URL, http.DefaultClient)
	header := make(http.Header)
	header.Set("X-Sequence-Project", strconv.Itoa(int(tenant.ProjectID)))
	ctx, err := proto.WithHTTPRequestHeaders(context.Background(), header)

	res, err := c.SendIntent(ctx, intent)
	require.NoError(t, err)
	assert.Equal(t, proto.IntentResponseCode("transactionReceipt"), res.Code)
	assert.NotEmpty(t, res.Data)
}
