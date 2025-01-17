package rpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/0xsequence/nitrocontrol/enclave"
	waasauthenticator "github.com/0xsequence/waas-authenticator"
	"github.com/0xsequence/waas-authenticator/config"
	"github.com/0xsequence/waas-authenticator/data"
	"github.com/0xsequence/waas-authenticator/proto"
	proto_wallet "github.com/0xsequence/waas-authenticator/proto/waas"
	"github.com/0xsequence/waas-authenticator/rpc/access"
	"github.com/0xsequence/waas-authenticator/rpc/attestation"
	"github.com/0xsequence/waas-authenticator/rpc/awscreds"
	"github.com/0xsequence/waas-authenticator/rpc/identity"
	"github.com/0xsequence/waas-authenticator/rpc/tenant"
	"github.com/0xsequence/waas-authenticator/rpc/tracing"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/go-chi/telemetry"
	"github.com/go-chi/traceid"
	"github.com/goware/cachestore/memlru"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
}

type RPC struct {
	Config     *config.Config
	Log        zerolog.Logger
	Server     *http.Server
	HTTPClient HTTPClient
	Enclave    *enclave.Enclave
	Tenants    *data.TenantTable
	Sessions   *data.SessionTable
	Accounts   *data.AccountTable
	Wallets    proto_wallet.WaaS
	Verifier   *identity.Verifier

	measurements *enclave.Measurements
	startTime    time.Time
	running      int32
}

func New(cfg *config.Config, client *http.Client) (*RPC, error) {
	if client == nil {
		client = http.DefaultClient
	}
	wrappedClient := tracing.WrapClient(client)

	options := []func(options *awsconfig.LoadOptions) error{
		awsconfig.WithRegion(cfg.Region),
		awsconfig.WithHTTPClient(wrappedClient),
		awsconfig.WithCredentialsProvider(awscreds.NewProvider(wrappedClient, cfg.Endpoints.MetadataServer)),
	}

	if cfg.Endpoints.AWSEndpoint != "" {
		options = append(options, awsconfig.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: cfg.Endpoints.AWSEndpoint}, nil
			}),
		), awsconfig.WithCredentialsProvider(&awscreds.StaticProvider{
			AccessKeyID:     "test",
			SecretAccessKey: "test",
			SessionToken:    "test",
		}))
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(), options...)
	if err != nil {
		return nil, err
	}

	tp, err := newOtelTracerProvider(context.Background(), client, cfg.Tracing)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{})

	// instrument all aws clients
	otelaws.AppendMiddlewares(&awsCfg.APIOptions)

	httpServer := &http.Server{
		ReadTimeout:       45 * time.Second,
		WriteTimeout:      45 * time.Second,
		IdleTimeout:       45 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	enclaveProvider := enclave.DummyProvider
	if cfg.Service.UseNSM {
		enclaveProvider = enclave.NitroProvider
	}
	enc, err := enclave.New(context.Background(), enclaveProvider, kms.NewFromConfig(awsCfg))
	if err != nil {
		return nil, err
	}

	m, err := enc.GetMeasurements(context.Background())
	if err != nil {
		return nil, err
	}

	cacheBackend := memlru.Backend(1024)
	verifier, err := identity.NewVerifier(cacheBackend, client)
	if err != nil {
		return nil, err
	}

	db := dynamodb.NewFromConfig(awsCfg)
	s := &RPC{
		Log: httplog.NewLogger("waas-authenticator", httplog.Options{
			LogLevel: zerolog.LevelDebugValue,
		}),
		Config:     cfg,
		Server:     httpServer,
		HTTPClient: wrappedClient,
		Enclave:    enc,
		Tenants:    data.NewTenantTable(db, cfg.Database.TenantsTable),
		Sessions:   data.NewSessionTable(db, cfg.Database.SessionsTable, "UserID-Index"),
		Accounts: data.NewAccountTable(db, cfg.Database.AccountsTable, data.AccountIndices{
			ByUserID: "UserID-Index",
			ByEmail:  "Email-Index",
		}),
		Wallets:      proto_wallet.NewWaaSClient(cfg.Endpoints.WaasAPIServer, wrappedClient),
		Verifier:     verifier,
		startTime:    time.Now(),
		measurements: m,
	}
	return s, nil
}

func (s *RPC) Run(ctx context.Context, l net.Listener) error {
	if s.IsRunning() {
		return fmt.Errorf("rpc: already running")
	}

	s.Log.Info().
		Str("op", "run").
		Str("ver", waasauthenticator.VERSION).
		Msgf("-> rpc: started enclave")

	atomic.StoreInt32(&s.running, 1)
	defer atomic.StoreInt32(&s.running, 0)

	// Setup HTTP server handler
	s.Server.Handler = s.Handler()

	// Handle stop signal to ensure clean shutdown
	go func() {
		<-ctx.Done()
		s.Stop(context.Background())
	}()

	// Start the http server and serve!
	err := s.Server.Serve(l)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *RPC) Stop(timeoutCtx context.Context) {
	if !s.IsRunning() || s.IsStopping() {
		return
	}
	atomic.StoreInt32(&s.running, 2)

	s.Log.Info().Str("op", "stop").Msg("-> rpc: stopping..")
	s.Server.Shutdown(timeoutCtx)
	s.Log.Info().Str("op", "stop").Msg("-> rpc: stopped.")
}

func (s *RPC) IsRunning() bool {
	return atomic.LoadInt32(&s.running) == 1
}

func (s *RPC) IsStopping() bool {
	return atomic.LoadInt32(&s.running) == 2
}

func (s *RPC) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)

	// Metrics and heartbeat
	r.Use(telemetry.Collector(s.Config.Telemetry, []string{"/rpc"}))
	r.Use(middleware.NoCache)
	r.Use(middleware.Heartbeat("/ping"))

	// Propagate TraceId
	r.Use(traceid.Middleware)

	// HTTP request logger
	r.Use(httplog.RequestLogger(s.Log, []string{"/", "/ping", "/status", "/favicon.ico"}))

	// Timeout any request after 28 seconds as Cloudflare has a 30 second limit anyways.
	r.Use(middleware.Timeout(28 * time.Second))

	// Quick pages
	r.Use(middleware.PageRoute("/", http.HandlerFunc(indexHandler)))
	r.Use(middleware.PageRoute("/status", http.HandlerFunc(s.statusHandler)))
	r.Use(middleware.PageRoute("/favicon.ico", http.HandlerFunc(emptyHandler)))

	// OpenTelemetry tracing
	r.Use(tracing.Middleware())

	// Generate attestation document
	r.Use(attestation.Middleware(s.Enclave))

	// Healthcheck
	r.Use(middleware.PageRoute("/health", http.HandlerFunc(s.healthHandler)))

	userRouter := r.Group(func(r chi.Router) {
		// Find and decrypt tenant data
		r.Use(tenant.Middleware(s.Tenants, s.Config.KMS.TenantKeys))
	})
	userRouter.Handle("/rpc/WaasAuthenticator/*", proto.NewWaasAuthenticatorServer(s))

	adminRouter := r.Group(func(r chi.Router) {
		// Validate admin JWTs
		r.Use(access.JWTAuthMiddleware(s.Config.Admin))
	})
	adminRouter.Handle("/rpc/WaasAuthenticatorAdmin/*", proto.NewWaasAuthenticatorAdminServer(s))

	if s.Config.Service.DebugProfiler {
		r.Mount("/debug", middleware.Profiler())
	}

	return r
}

// Ping is a healthcheck that returns an empty message.
func (s *RPC) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("."))
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func newOtelTracerProvider(ctx context.Context, client *http.Client, cfg config.TracingConfig) (*trace.TracerProvider, error) {
	traceExporter, err := zipkin.New(cfg.Endpoint, zipkin.WithClient(client))
	if err != nil {
		return nil, fmt.Errorf("failed to create new OTLP trace exporter: %v", err)
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("WaasAuthenticator"),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.ParentBased(tracing.SampleRoutes("/rpc/WaasAuthenticator/*", "/rpc/WaasAuthenticatorAdmin/*"))),
		trace.WithBatcher(traceExporter),
		trace.WithIDGenerator(xray.NewIDGenerator()),
		trace.WithResource(r),
	)
	return tp, nil
}
