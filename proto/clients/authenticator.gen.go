// sequence-waas-authenticator v0.1.0 f7bf79363b992ace30e01ed2c6700c9ad0336c95
// --
// Code generated by webrpc-gen@v0.18.6 with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=authenticator.ridl -target=golang -pkg=proto -client -out=./clients/authenticator.gen.go
package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/goware/validation"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.1.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "f7bf79363b992ace30e01ed2c6700c9ad0336c95"
}

//
// Common types
//

type IntentName string

const (
	IntentName_openSession           IntentName = "openSession"
	IntentName_closeSession          IntentName = "closeSession"
	IntentName_validateSession       IntentName = "validateSession"
	IntentName_finishValidateSession IntentName = "finishValidateSession"
	IntentName_listSessions          IntentName = "listSessions"
	IntentName_getSession            IntentName = "getSession"
	IntentName_sessionAuthProof      IntentName = "sessionAuthProof"
	IntentName_feeOptions            IntentName = "feeOptions"
	IntentName_signMessage           IntentName = "signMessage"
	IntentName_sendTransaction       IntentName = "sendTransaction"
	IntentName_getTransactionReceipt IntentName = "getTransactionReceipt"
	IntentName_federateAccount       IntentName = "federateAccount"
	IntentName_removeAccount         IntentName = "removeAccount"
	IntentName_listAccounts          IntentName = "listAccounts"
)

func (x IntentName) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *IntentName) UnmarshalText(b []byte) error {
	*x = IntentName(string(b))
	return nil
}

func (x *IntentName) Is(values ...IntentName) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type IntentResponseCode string

const (
	IntentResponseCode_sessionOpened      IntentResponseCode = "sessionOpened"
	IntentResponseCode_sessionClosed      IntentResponseCode = "sessionClosed"
	IntentResponseCode_sessionList        IntentResponseCode = "sessionList"
	IntentResponseCode_validationRequired IntentResponseCode = "validationRequired"
	IntentResponseCode_validationStarted  IntentResponseCode = "validationStarted"
	IntentResponseCode_validationFinished IntentResponseCode = "validationFinished"
	IntentResponseCode_sessionAuthProof   IntentResponseCode = "sessionAuthProof"
	IntentResponseCode_signedMessage      IntentResponseCode = "signedMessage"
	IntentResponseCode_feeOptions         IntentResponseCode = "feeOptions"
	IntentResponseCode_transactionReceipt IntentResponseCode = "transactionReceipt"
	IntentResponseCode_transactionFailed  IntentResponseCode = "transactionFailed"
	IntentResponseCode_getSessionResponse IntentResponseCode = "getSessionResponse"
	IntentResponseCode_accountList        IntentResponseCode = "accountList"
	IntentResponseCode_accountFederated   IntentResponseCode = "accountFederated"
	IntentResponseCode_accountRemoved     IntentResponseCode = "accountRemoved"
)

func (x IntentResponseCode) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *IntentResponseCode) UnmarshalText(b []byte) error {
	*x = IntentResponseCode(string(b))
	return nil
}

func (x *IntentResponseCode) Is(values ...IntentResponseCode) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type Intent struct {
	Version    string       `json:"version"`
	Name       IntentName   `json:"name"`
	ExpiresAt  uint64       `json:"expiresAt"`
	IssuedAt   uint64       `json:"issuedAt"`
	Data       interface{}  `json:"data"`
	Signatures []*Signature `json:"signatures,omitempty"`
}

type Signature struct {
	SessionID string `json:"sessionId"`
	Signature string `json:"signature"`
}

type IntentResponse struct {
	Code IntentResponseCode `json:"code"`
	Data interface{}        `json:"data"`
}

type IdentityType uint8

const (
	IdentityType_None  IdentityType = 0
	IdentityType_Guest IdentityType = 1
	IdentityType_OIDC  IdentityType = 2
)

var IdentityType_name = map[uint8]string{
	0: "None",
	1: "Guest",
	2: "OIDC",
}

var IdentityType_value = map[string]uint8{
	"None":  0,
	"Guest": 1,
	"OIDC":  2,
}

func (x IdentityType) String() string {
	return IdentityType_name[uint8(x)]
}

func (x IdentityType) MarshalText() ([]byte, error) {
	return []byte(IdentityType_name[uint8(x)]), nil
}

func (x *IdentityType) UnmarshalText(b []byte) error {
	*x = IdentityType(IdentityType_value[string(b)])
	return nil
}

func (x *IdentityType) Is(values ...IdentityType) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	AppVersion    string `json:"appVersion"`
}

type RuntimeStatus struct {
	// overall status, true/false
	HealthOK  bool      `json:"healthOK"`
	StartTime time.Time `json:"startTime"`
	Uptime    uint64    `json:"uptime"`
	Ver       string    `json:"ver"`
	PCR0      string    `json:"pcr0"`
}

type Chain struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	IsEnabled bool   `json:"isEnabled"`
}

type Identity struct {
	Type    IdentityType `json:"type"`
	Issuer  string       `json:"iss"`
	Subject string       `json:"sub"`
	Email   string       `json:"email,omitempty"`
}

type OpenIdProvider struct {
	Issuer   string   `json:"iss"`
	Audience []string `json:"aud"`
}

type Tenant struct {
	ProjectID      uint64             `json:"projectId"`
	Version        int                `json:"version"`
	OIDCProviders  []*OpenIdProvider  `json:"oidcProviders"`
	AllowedOrigins validation.Origins `json:"allowedOrigins"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}

type TenantData struct {
	ProjectID       uint64               `json:"projectId"`
	PrivateKey      string               `json:"privateKey"`
	ParentAddress   common.Address       `json:"parentAddress"`
	UserSalt        hexutil.Bytes        `json:"userSalt"`
	SequenceContext *MiniSequenceContext `json:"sequenceContext"`
	UpgradeCode     string               `json:"upgradeCode"`
	WaasAccessToken string               `json:"waasAccessToken"`
	OIDCProviders   []*OpenIdProvider    `json:"oidcProviders"`
	KMSKeys         []string             `json:"kmsKeys"`
	AllowedOrigins  validation.Origins   `json:"allowedOrigins"`
}

type MiniSequenceContext struct {
	Factory    string `json:"factory"`
	MainModule string `json:"mainModule"`
}

type AccountData struct {
	ProjectID uint64    `json:"projectId"`
	UserID    string    `json:"userId"`
	Identity  string    `json:"identity"`
	CreatedAt time.Time `json:"createdAt"`
}

type Session struct {
	ID           string    `json:"id"`
	ProjectID    uint64    `json:"projectId"`
	UserID       string    `json:"userId"`
	Identity     Identity  `json:"identity"`
	FriendlyName string    `json:"friendlyName"`
	CreatedAt    time.Time `json:"createdAt"`
	RefreshedAt  time.Time `json:"refreshedAt"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

type SessionData struct {
	ID        string    `json:"id"`
	ProjectID uint64    `json:"projectId"`
	UserID    string    `json:"userId"`
	Identity  string    `json:"identity"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

var WebRPCServices = map[string][]string{
	"WaasAuthenticator": {
		"RegisterSession",
		"SendIntent",
		"ChainList",
	},
	"WaasAuthenticatorAdmin": {
		"Version",
		"RuntimeStatus",
		"Clock",
		"GetTenant",
		"CreateTenant",
		"UpdateTenant",
	},
}

//
// Server types
//

type WaasAuthenticator interface {
	RegisterSession(ctx context.Context, intent *Intent, friendlyName string) (*Session, *IntentResponse, error)
	SendIntent(ctx context.Context, intent *Intent) (*IntentResponse, error)
	ChainList(ctx context.Context) ([]*Chain, error)
}

type WaasAuthenticatorAdmin interface {
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	Clock(ctx context.Context) (time.Time, error)
	GetTenant(ctx context.Context, projectId uint64) (*Tenant, error)
	CreateTenant(ctx context.Context, projectId uint64, waasAccessToken string, oidcProviders []*OpenIdProvider, allowedOrigins []string, password *string) (*Tenant, string, error)
	UpdateTenant(ctx context.Context, projectId uint64, upgradeCode string, oidcProviders []*OpenIdProvider, allowedOrigins []string) (*Tenant, error)
}

//
// Client types
//

type WaasAuthenticatorClient interface {
	RegisterSession(ctx context.Context, intent *Intent, friendlyName string) (*Session, *IntentResponse, error)
	SendIntent(ctx context.Context, intent *Intent) (*IntentResponse, error)
	ChainList(ctx context.Context) ([]*Chain, error)
}

type WaasAuthenticatorAdminClient interface {
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	Clock(ctx context.Context) (time.Time, error)
	GetTenant(ctx context.Context, projectId uint64) (*Tenant, error)
	CreateTenant(ctx context.Context, projectId uint64, waasAccessToken string, oidcProviders []*OpenIdProvider, allowedOrigins []string, password *string) (*Tenant, string, error)
	UpdateTenant(ctx context.Context, projectId uint64, upgradeCode string, oidcProviders []*OpenIdProvider, allowedOrigins []string) (*Tenant, error)
}

//
// Client
//

const WaasAuthenticatorPathPrefix = "/rpc/WaasAuthenticator/"
const WaasAuthenticatorAdminPathPrefix = "/rpc/WaasAuthenticatorAdmin/"

type waasAuthenticatorClient struct {
	client HTTPClient
	urls   [3]string
}

func NewWaasAuthenticatorClient(addr string, client HTTPClient) WaasAuthenticatorClient {
	prefix := urlBase(addr) + WaasAuthenticatorPathPrefix
	urls := [3]string{
		prefix + "RegisterSession",
		prefix + "SendIntent",
		prefix + "ChainList",
	}
	return &waasAuthenticatorClient{
		client: client,
		urls:   urls,
	}
}

func (c *waasAuthenticatorClient) RegisterSession(ctx context.Context, intent *Intent, friendlyName string) (*Session, *IntentResponse, error) {
	in := struct {
		Arg0 *Intent `json:"intent"`
		Arg1 string  `json:"friendlyName"`
	}{intent, friendlyName}
	out := struct {
		Ret0 *Session        `json:"session"`
		Ret1 *IntentResponse `json:"response"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *waasAuthenticatorClient) SendIntent(ctx context.Context, intent *Intent) (*IntentResponse, error) {
	in := struct {
		Arg0 *Intent `json:"intent"`
	}{intent}
	out := struct {
		Ret0 *IntentResponse `json:"response"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *waasAuthenticatorClient) ChainList(ctx context.Context) ([]*Chain, error) {
	out := struct {
		Ret0 []*Chain `json:"chains"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

type waasAuthenticatorAdminClient struct {
	client HTTPClient
	urls   [6]string
}

func NewWaasAuthenticatorAdminClient(addr string, client HTTPClient) WaasAuthenticatorAdminClient {
	prefix := urlBase(addr) + WaasAuthenticatorAdminPathPrefix
	urls := [6]string{
		prefix + "Version",
		prefix + "RuntimeStatus",
		prefix + "Clock",
		prefix + "GetTenant",
		prefix + "CreateTenant",
		prefix + "UpdateTenant",
	}
	return &waasAuthenticatorAdminClient{
		client: client,
		urls:   urls,
	}
}

func (c *waasAuthenticatorAdminClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *waasAuthenticatorAdminClient) RuntimeStatus(ctx context.Context) (*RuntimeStatus, error) {
	out := struct {
		Ret0 *RuntimeStatus `json:"status"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *waasAuthenticatorAdminClient) Clock(ctx context.Context) (time.Time, error) {
	out := struct {
		Ret0 time.Time `json:"serverTime"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *waasAuthenticatorAdminClient) GetTenant(ctx context.Context, projectId uint64) (*Tenant, error) {
	in := struct {
		Arg0 uint64 `json:"projectId"`
	}{projectId}
	out := struct {
		Ret0 *Tenant `json:"tenant"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[3], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

func (c *waasAuthenticatorAdminClient) CreateTenant(ctx context.Context, projectId uint64, waasAccessToken string, oidcProviders []*OpenIdProvider, allowedOrigins []string, password *string) (*Tenant, string, error) {
	in := struct {
		Arg0 uint64            `json:"projectId"`
		Arg1 string            `json:"waasAccessToken"`
		Arg2 []*OpenIdProvider `json:"oidcProviders"`
		Arg3 []string          `json:"allowedOrigins"`
		Arg4 *string           `json:"password"`
	}{projectId, waasAccessToken, oidcProviders, allowedOrigins, password}
	out := struct {
		Ret0 *Tenant `json:"tenant"`
		Ret1 string  `json:"upgradeCode"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[4], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *waasAuthenticatorAdminClient) UpdateTenant(ctx context.Context, projectId uint64, upgradeCode string, oidcProviders []*OpenIdProvider, allowedOrigins []string) (*Tenant, error) {
	in := struct {
		Arg0 uint64            `json:"projectId"`
		Arg1 string            `json:"upgradeCode"`
		Arg2 []*OpenIdProvider `json:"oidcProviders"`
		Arg3 []string          `json:"allowedOrigins"`
	}{projectId, upgradeCode, oidcProviders, allowedOrigins}
	out := struct {
		Ret0 *Tenant `json:"tenant"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[5], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to close response body: %w", cerr))
		}
	}

	return out.Ret0, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doHTTPRequest is common code to make a request to the remote service.
func doHTTPRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(fmt.Errorf("could not build request: %w", err))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(err)
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return nil, rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCause(fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return resp, nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}
	HTTPRequestCtxKey              = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func ServiceNameFromContext(ctx context.Context) string {
	service, _ := ctx.Value(ServiceNameCtxKey).(string)
	return service
}

func MethodNameFromContext(ctx context.Context) string {
	method, _ := ctx.Value(MethodNameCtxKey).(string)
	return method
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	return r
}

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if target == nil {
		return false
	}
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func (e WebRPCError) WithCause(cause error) WebRPCError {
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

func (e WebRPCError) WithCausef(format string, args ...interface{}) WebRPCError {
	cause := fmt.Errorf(format, args...)
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Deprecated: Use .WithCause() method on WebRPCError.
func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	return rpcErr.WithCause(cause)
}

// Webrpc errors
var (
	ErrWebrpcEndpoint           = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed      = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute           = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod          = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest         = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse        = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic        = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError      = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
	ErrWebrpcClientDisconnected = WebRPCError{Code: -8, Name: "WebrpcClientDisconnected", Message: "client disconnected", HTTPStatus: 400}
	ErrWebrpcStreamLost         = WebRPCError{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400}
	ErrWebrpcStreamFinished     = WebRPCError{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200}
)

// Schema errors
var (
	ErrUnauthorized   = WebRPCError{Code: 1000, Name: "Unauthorized", Message: "Unauthorized access", HTTPStatus: 401}
	ErrTenantNotFound = WebRPCError{Code: 1001, Name: "TenantNotFound", Message: "Tenant not found", HTTPStatus: 404}
)
