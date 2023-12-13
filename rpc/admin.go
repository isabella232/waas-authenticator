package rpc

import (
	"context"
	"encoding/base32"
	"fmt"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"golang.org/x/sync/errgroup"

	"github.com/0xsequence/waas-authenticator/data"
	"github.com/0xsequence/waas-authenticator/proto"
	"github.com/0xsequence/waas-authenticator/rpc/attestation"
	"github.com/0xsequence/waas-authenticator/rpc/crypto"
)

func (s *RPC) GetTenant(ctx context.Context, projectID uint64) (*proto.Tenant, error) {
	tnt, _, err := s.Tenants.GetLatest(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if tnt == nil {
		return nil, fmt.Errorf("tenant not found")
	}

	tenantData, _, err := crypto.DecryptData[*proto.TenantData](ctx, tnt.EncryptedKey, tnt.Ciphertext, s.Config.KMS.TenantKeys)
	if err != nil {
		return nil, fmt.Errorf("decrypt tenant data: %w", err)
	}

	retTenant := &proto.Tenant{
		ProjectID:     tnt.ProjectID,
		Version:       tnt.Version,
		OIDCProviders: tenantData.OIDCProviders,
		UpdatedAt:     tnt.CreatedAt,
	}
	return retTenant, nil
}

func (s *RPC) CreateTenant(
	ctx context.Context, projectID uint64, waasAccessToken string, oidcProviders []*proto.OpenIdProvider,
) (*proto.Tenant, string, error) {
	att := attestation.FromContext(ctx)

	_, found, err := s.Tenants.GetLatest(ctx, projectID)
	if err != nil {
		return nil, "", err
	}
	if found {
		return nil, "", fmt.Errorf("tenant already exists")
	}

	if err := validateOIDCProviders(ctx, s.HTTPClient, oidcProviders); err != nil {
		return nil, "", fmt.Errorf("invalid oidcProviders: %w", err)
	}

	wallet, err := ethwallet.NewWalletFromRandomEntropy()
	if err != nil {
		return nil, "", fmt.Errorf("generating wallet: %w", err)
	}

	waasCtx, err := waasContext(ctx, waasAccessToken)
	if err != nil {
		return nil, "", err
	}

	// TODO: these are 4 calls to WaaS API, can we do it all in one call?
	if _, err := s.Wallets.UseHotWallet(waasCtx, wallet.Address().String()); err != nil {
		return nil, "", err
	}

	userSalt, err := s.Wallets.UserSalt(waasCtx)
	if err != nil {
		return nil, "", err
	}
	userSaltBytes, err := hexutil.Decode(userSalt)
	if err != nil {
		return nil, "", err
	}

	parentAddress, err := s.Wallets.ParentWallet(waasCtx)
	if err != nil {
		return nil, "", err
	}

	seqContext, err := s.Wallets.SequenceContext(waasCtx)
	if err != nil {
		return nil, "", err
	}

	upgradeCode := make([]byte, 10)
	if _, err := att.Read(upgradeCode); err != nil {
		return nil, "", err
	}

	privateKey := wallet.PrivateKeyHex()[2:] // remove 0x prefix
	tenantData := proto.TenantData{
		ProjectID:     projectID,
		PrivateKey:    privateKey,
		ParentAddress: common.HexToAddress(parentAddress),
		UserSalt:      userSaltBytes,
		SequenceContext: &proto.MiniSequenceContext{
			Factory:    seqContext.Factory,
			MainModule: seqContext.MainModule,
		},
		UpgradeCode:     base32.StdEncoding.EncodeToString(upgradeCode),
		WaasAccessToken: waasAccessToken,
		OIDCProviders:   oidcProviders,
		TransportKeys:   s.Config.KMS.DefaultTransportKeys,
		SessionKeys:     s.Config.KMS.DefaultSessionKeys,
	}

	encryptedKey, algorithm, ciphertext, err := crypto.EncryptData(ctx, att, s.Config.KMS.TenantKeys[0], tenantData)
	if err != nil {
		return nil, "", fmt.Errorf("encrypting tenant data: %w", err)
	}

	dbTenant := &data.Tenant{
		ProjectID:    projectID,
		Version:      1,
		EncryptedKey: encryptedKey,
		Algorithm:    algorithm,
		Ciphertext:   ciphertext,
		CreatedAt:    time.Now(),
	}
	if err := s.Tenants.Add(ctx, dbTenant); err != nil {
		return nil, "", err
	}

	retTenant := &proto.Tenant{
		ProjectID:     projectID,
		Version:       dbTenant.Version,
		OIDCProviders: tenantData.OIDCProviders,
		UpdatedAt:     dbTenant.CreatedAt,
	}
	return retTenant, tenantData.UpgradeCode, nil
}

func (s *RPC) UpdateTenant(ctx context.Context, projectID uint64, oidcProviders []*proto.OpenIdProvider) (*proto.Tenant, error) {
	// TODO: should validate the upgrade code!
	att := attestation.FromContext(ctx)

	tnt, found, err := s.Tenants.GetLatest(ctx, projectID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("tenant not found")
	}

	if err := validateOIDCProviders(ctx, s.HTTPClient, oidcProviders); err != nil {
		return nil, fmt.Errorf("invalid oidcProviders: %w", err)
	}

	tntData, _, err := crypto.DecryptData[*proto.TenantData](ctx, tnt.EncryptedKey, tnt.Ciphertext, s.Config.KMS.TenantKeys)
	if err != nil {
		return nil, fmt.Errorf("decrypt tenant data: %w", err)
	}

	tntData.OIDCProviders = oidcProviders

	encryptedKey, algorithm, ciphertext, err := crypto.EncryptData(ctx, att, s.Config.KMS.TenantKeys[0], tntData)
	if err != nil {
		return nil, fmt.Errorf("encrypting tenant data: %w", err)
	}

	tnt.Version += 1
	tnt.EncryptedKey = encryptedKey
	tnt.Ciphertext = ciphertext
	tnt.Algorithm = algorithm

	if err := s.Tenants.Add(ctx, tnt); err != nil {
		return nil, err
	}

	retTenant := &proto.Tenant{
		ProjectID:     tnt.ProjectID,
		Version:       tnt.Version,
		OIDCProviders: tntData.OIDCProviders,
		UpdatedAt:     tnt.CreatedAt,
	}
	return retTenant, nil
}

func validateOIDCProviders(ctx context.Context, client HTTPClient, providers []*proto.OpenIdProvider) error {
	var wg errgroup.Group
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i, provider := range providers {
		provider := provider

		if provider.Issuer == "" {
			return fmt.Errorf("provider %d: empty issuer", i)
		}

		if provider.Audience == nil {
			return fmt.Errorf("provider %d: audience is required", i)
		}

		wg.Go(func() error {
			if _, err := getProviderKeySet(ctx, client, provider.Issuer); err != nil {
				return err
			}
			return nil
		})
	}

	return wg.Wait()
}