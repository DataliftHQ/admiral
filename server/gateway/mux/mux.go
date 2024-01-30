package mux

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	"net/textproto"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	gatewayv1 "go.datalift.io/admiral/server/config/gateway/v1"
)

const (
	xHeader        = "X-"
	xForwardedFor  = "X-Forwarded-For"
	xForwardedHost = "X-Forwarded-Host"
)

var apiPattern = regexp.MustCompile(`^/api/v\d+/`)

type assetHandler struct {
	assetCfg *gatewayv1.Assets

	next       http.Handler
	fileSystem http.FileSystem
	fileServer http.Handler
}

func copyHTTPResponse(resp *http.Response, w http.ResponseWriter) {
	for key, values := range resp.Header {
		for _, val := range values {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func (a *assetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if apiPattern.MatchString(r.URL.Path) || r.URL.Path == "/healthz" {
		a.next.ServeHTTP(w, r)
		return
	}

	a.fileServer.ServeHTTP(w, r)
}

func newCustomResponseForwarder(secureCookies bool) func(context.Context, http.ResponseWriter, proto.Message) error {
	return func(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
		md, ok := runtime.ServerMetadataFromContext(ctx)
		if !ok {
			return nil
		}

		if cookies := md.HeaderMD.Get("Set-Cookie-Token"); len(cookies) > 0 {
			cookie := &http.Cookie{
				Name:     "token",
				Value:    cookies[0],
				Path:     "/",
				HttpOnly: false,
				Secure:   secureCookies,
			}
			http.SetCookie(w, cookie)
		}

		if cookies := md.HeaderMD.Get("Set-Cookie-Refresh-Token"); len(cookies) > 0 {
			cookie := &http.Cookie{
				Name:     "refreshToken",
				Value:    cookies[0],
				Path:     "/api/v1/authn/login",
				HttpOnly: true, // Client cannot access refresh token, it is sent by browser only if login is attempted.
				Secure:   secureCookies,
			}
			http.SetCookie(w, cookie)
		}

		// Redirect if it's the browser (non-XHR).
		redirects := md.HeaderMD.Get("Location")
		if len(redirects) > 0 && isBrowser(requestHeadersFromResponseWriter(w)) {
			code := http.StatusFound
			if st := md.HeaderMD.Get("Location-Status"); len(st) > 0 {
				headerCodeOverride, err := strconv.Atoi(st[0])
				if err != nil {
					return err
				}
				code = headerCodeOverride
			}

			w.Header().Set("Location", redirects[0])
			w.WriteHeader(code)
		}

		return nil
	}
}

func customHeaderMatcher(key string) (string, bool) {
	key = textproto.CanonicalMIMEHeaderKey(key)
	if strings.HasPrefix(key, xHeader) {
		// exclude handling these headers as they are looked up by grpc's annotate context flow and added to the context
		// metadata if they're not found
		if key != xForwardedFor && key != xForwardedHost {
			return runtime.MetadataPrefix + key, true
		}
	}
	// the the default header mapping rule
	return runtime.DefaultHeaderMatcher(key)
}

func customErrorHandler(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	if isBrowser(req.Header) { // Redirect if it's the browser (non-XHR).
		if s, ok := status.FromError(err); ok && s.Code() == codes.Unauthenticated {
			redirectPath := fmt.Sprintf("/api/v1/authn/login?redirect_url=%s", url.QueryEscape(req.RequestURI))
			http.Redirect(w, req, redirectPath, http.StatusFound)
			return
		}
	}

	runtime.DefaultHTTPErrorHandler(ctx, mux, m, w, req, err)
}

func New(unaryInterceptors []grpc.UnaryServerInterceptor, assets http.FileSystem, metricsHandler http.Handler, gatewayCfg *gatewayv1.GatewayOptions) (*Mux, error) {
	secureCookies := true
	if gatewayCfg.SecureCookies != nil {
		secureCookies = gatewayCfg.SecureCookies.Value
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(unaryInterceptors...))
	jsonGateway := runtime.NewServeMux(
		runtime.WithForwardResponseOption(newCustomResponseForwarder(secureCookies)),
		runtime.WithErrorHandler(customErrorHandler),
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					// Use camelCase for the JSON version.
					UseProtoNames: false,
					// Transmit zero-values over the wire.
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{},
			},
		),
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
	)

	// If there is a configured asset provider, we check to see if the service is configured before proceeding.
	// Bailing out early during the startup process instead of hitting this error at runtime when serving assets.
	//if gatewayCfg.Assets != nil && gatewayCfg.Assets.Provider != nil {
	//	_, err := getAssetProviderService(gatewayCfg.Assets)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", &assetHandler{
		assetCfg:   gatewayCfg.Assets,
		next:       jsonGateway,
		fileSystem: assets,
		fileServer: http.FileServer(assets),
	})

	if gatewayCfg.EnablePprof {
		httpMux.HandleFunc("/debug/pprof/", pprof.Index)
	}

	if metricsHandler != nil {
		_, ok := gatewayCfg.Stats.Reporter.(*gatewayv1.Stats_PrometheusReporter_)
		if !ok {
			return nil, fmt.Errorf("Expected *gatewayv1.Stats_PrometheusReporter_, got %T", gatewayCfg.Stats.Reporter)
		}
		metricsPath := "/metrics"
		promCfg := gatewayCfg.Stats.GetPrometheusReporter()
		if promCfg.HandlerPath != "" {
			metricsPath = promCfg.HandlerPath
		}
		httpMux.Handle(metricsPath, metricsHandler)
	}

	mux := &Mux{
		GRPCServer:  grpcServer,
		JSONGateway: jsonGateway,
		HTTPMux:     httpMux,
	}
	return mux, nil
}

// Mux allows sharing one port between gRPC and the corresponding JSON gateway via header-based multiplexing.
type Mux struct {
	// Create empty handlers for gRPC and grpc-gateway (JSON) traffic.
	JSONGateway *runtime.ServeMux
	HTTPMux     http.Handler
	GRPCServer  *grpc.Server
}

// Adapted from https://github.com/grpc/grpc-go/blob/197c621/server.go#L760-L778.
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
		m.GRPCServer.ServeHTTP(w, r)
	} else {
		m.HTTPMux.ServeHTTP(w, r)
	}
}

func (m *Mux) EnableGRPCReflection() {
	reflection.Register(m.GRPCServer)
}

// "h2c" is the unencrypted form of HTTP/2.
func InsecureHandler(handler http.Handler) http.Handler {
	return h2c.NewHandler(handler, &http2.Server{})
}
