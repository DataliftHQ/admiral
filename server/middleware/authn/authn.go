package authn

import (
	"context"
	"errors"
	"strings"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"go.datalift.io/admiral/server/gateway/mux"
	"go.datalift.io/admiral/server/middleware"
	"go.datalift.io/admiral/server/service"
	"go.datalift.io/admiral/server/service/authn"
)

const Name = "admiral.middleware.authn"

func New(cfg *any.Any, logger *zap.Logger, scope tally.Scope) (middleware.Middleware, error) {
	svc, ok := service.Registry["admiral.service.authn"]
	if !ok {
		return nil, errors.New("unable to get authn service")
	}

	p, ok := svc.(authn.Provider)
	if !ok {
		return nil, errors.New("authn service was not the correct type")
	}

	m := &mid{
		provider: p,
	}

	return m, nil
}

type mid struct {
	provider authn.Provider
}

func (m *mid) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check for auth.
		authenticatedCtx, authErr := m.authenticate(ctx)

		// Determine if it's on the allow list.
		checkRequired := true
		for _, allow := range authn.AlwaysAllowedMethods {
			if middleware.MatchMethodOrResource(allow, info.FullMethod) {
				checkRequired = false
				break
			}
		}

		// Assert auth if required.
		if checkRequired {
			if authErr != nil {
				return nil, status.New(codes.Unauthenticated, authErr.Error()).Err()
			}
			return handler(authenticatedCtx, req)
		}

		// If auth not required, we still append claims for logging purposes or anonymously accessible APIs.
		if _, err := authn.ClaimsFromContext(authenticatedCtx); err != nil {
			// Anonymous claims if there weren't any authenticated claims.
			return handler(authn.ContextWithAnonymousClaims(ctx), req)
		}
		return handler(authenticatedCtx, req)
	}
}

// getToken looks for the token in the authorization header or cookies.
func getToken(md metadata.MD) (string, error) {
	if tokens := md.Get("authorization"); len(tokens) > 0 {
		splitToken := strings.Split(tokens[0], "Bearer")
		if len(splitToken) != 2 {
			return "", errors.New("bad token format, expected Authorization: Bearer <token>")
		}
		return strings.TrimSpace(splitToken[1]), nil
	}

	v := md.Get("grpcgateway-cookie")
	if len(v) == 0 {
		return "", errors.New("token not present in authorization header or cookies")
	}

	return mux.GetCookieValue(v, "token")
}

func (m *mid) authenticate(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no headers present on request")
	}

	token, err := getToken(md)
	if err != nil {
		return nil, err
	}

	// Verify token.
	claims, err := m.provider.Verify(ctx, token)
	if err != nil {
		return nil, err
	}

	// Append claims information to context.
	return authn.ContextWithClaims(ctx, claims), nil
}
