package authn

import (
	"context"
	"errors"
	"fmt"
	"go.datalift.io/admiral/server/gateway/log"
	"go.datalift.io/admiral/server/gateway/mux"
	"go.datalift.io/admiral/server/service"
	"go.datalift.io/admiral/server/service/authn"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	authnv1 "go.datalift.io/admiral/common/api/authn/v1"
	"go.datalift.io/admiral/server/endpoint"
)

const (
	Name = "admiral.endpoint.authn"
)

type endp struct {
	provider authn.Provider
	issuer   authn.Issuer

	logger *zap.Logger
	scope  tally.Scope
}

func New(_ *anypb.Any, log *zap.Logger, scope tally.Scope) (endpoint.Endpoint, error) {
	svc, ok := service.Registry["admiral.service.authn"]
	if !ok {
		return nil, errors.New("unable to get authn service")
	}

	p, ok := svc.(authn.Service)
	if !ok {
		return nil, errors.New("authn service was not the correct type")
	}

	return &endp{
		provider: p,
		issuer:   p,
		logger:   log,
		scope:    scope,
	}, nil
}

func (e *endp) Register(r endpoint.Registrar) error {
	authnv1.RegisterAuthnAPIServer(r.GRPCServer(), e)
	return r.RegisterJSONGateway(authnv1.RegisterAuthnAPIHandler)
}

func (e *endp) loginViaRefresh(ctx context.Context, redirectURL string) (*authnv1.LoginResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil
	}

	cookies := md.Get("grpcgateway-cookie")
	if len(cookies) == 0 {
		return nil, nil
	}

	refreshToken, err := mux.GetCookieValue(cookies, "refreshToken")
	if err != nil {
		return nil, nil
	}

	token, err := e.issuer.RefreshToken(ctx, &oauth2.Token{RefreshToken: refreshToken})
	if err != nil {
		return nil, err
	}

	err = grpc.SetHeader(ctx, metadata.New(map[string]string{
		"Location":                 redirectURL,
		"Set-Cookie-Token":         token.AccessToken,
		"Set-Cookie-Refresh-Token": token.RefreshToken,
	}))
	if err != nil {
		return nil, err
	}

	return &authnv1.LoginResponse{
		Return: &authnv1.LoginResponse_Token_{
			Token: &authnv1.LoginResponse_Token{
				AccessToken:  token.AccessToken,
				RefreshToken: token.RefreshToken,
			},
		},
	}, nil
}

func (e *endp) Login(ctx context.Context, req *authnv1.LoginRequest) (*authnv1.LoginResponse, error) {
	resp, err := e.loginViaRefresh(ctx, req.RedirectUrl)
	if err != nil {
		e.logger.Info("login via refresh token failed, continuing regular auth flow", log.ErrorField(err))
	} else if resp != nil {
		return resp, nil
	}

	state, err := e.provider.GetStateNonce(req.RedirectUrl)
	if err != nil {
		return nil, err
	}
	authURL, err := e.provider.GetAuthCodeURL(ctx, state)
	if err != nil {
		return nil, err
	}

	if err := grpc.SetHeader(ctx, metadata.Pairs("Location", authURL)); err != nil {
		return nil, err
	}

	return &authnv1.LoginResponse{
		Return: &authnv1.LoginResponse_AuthUrl{AuthUrl: authURL},
	}, nil
}

func (e *endp) Callback(ctx context.Context, req *authnv1.CallbackRequest) (*authnv1.CallbackResponse, error) {
	if req.Error != "" {
		return nil, fmt.Errorf("%s: %s", req.Error, req.ErrorDescription)
	}

	redirectURL, err := e.provider.ValidateStateNonce(req.State)
	if err != nil {
		return nil, err
	}

	token, err := e.provider.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	md := metadata.New(map[string]string{
		"Location":         redirectURL,
		"Set-Cookie-Token": token.AccessToken,
	})

	if token.RefreshToken != "" {
		md.Set("Set-Cookie-Refresh-Token", token.RefreshToken)
	}

	if err := grpc.SetHeader(ctx, md); err != nil {
		return nil, err
	}

	return &authnv1.CallbackResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}

func (e *endp) CreateToken(ctx context.Context, req *authnv1.CreateTokenRequest) (*authnv1.CreateTokenResponse, error) {
	return nil, nil
}
