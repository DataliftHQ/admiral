package authn

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/uber-go/tally/v4"
	authnv1 "go.datalift.io/admiral/server/config/service/authn/v1"
	"go.datalift.io/admiral/server/service"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/types/known/anypb"
)

var AlwaysAllowedMethods = []string{
	"/admiral.authn.v1.AuthnAPI/Callback",
	"/admiral.authn.v1.AuthnAPI/Login",
	"/admiral.healthcheck.v1.HealthcheckAPI/*",
	"/admiral.settings.v1.SettingsAPI/*",
}

const Name = "admiral.service.authn"

func New(cfg *anypb.Any, _ *zap.Logger, _ tally.Scope) (service.Service, error) {
	config := &authnv1.Config{}
	if err := cfg.UnmarshalTo(config); err != nil {
		return nil, err
	}

	// DATABASE POSTGRES AS TOKEN STORAGE
	//db, _ := service.Registry[]

	tokenStorage, _ := service.Registry[StorageName].(Storage)

	switch t := config.Type.(type) {
	case *authnv1.Config_Oidc:
		return NewOIDCProvider(context.Background(), config, tokenStorage)
	default:
		return nil, fmt.Errorf("authn provider type '%T' not implemented", t)
	}
}

// Standardized representation of a user's claims.
type Claims struct {
	*jwt.RegisteredClaims

	// Groups could be derived from the token or an external mapping.
	Groups []string `json:"groups,omitempty"`
}

type Provider interface {
	GetStateNonce(redirectURL string) (string, error)
	ValidateStateNonce(state string) (redirectURL string, err error)

	Verify(ctx context.Context, rawIDToken string) (*Claims, error)
	GetAuthCodeURL(ctx context.Context, state string) (string, error)
	Exchange(ctx context.Context, code string) (token *oauth2.Token, err error)
}

type Issuer interface {
	// CreateToken creates a new OAuth2 for the provided subject with the provided expiration. If expiry is nil,
	// the token will never expire.
	//CreateToken(ctx context.Context, subject string, tokenType authnmodulev1.CreateTokenRequest_TokenType, expiry *time.Duration) (token *oauth2.Token, err error)
	RefreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error)
}

type Service interface {
	Issuer
	Provider
	TokenReader // Read calls are proxied through the IssuerProvider so the token can be refreshed if needed.
}

type TokenReader interface {
	Read(ctx context.Context, userID, provider string) (*oauth2.Token, error)
}

type TokenStorer interface {
	Store(ctx context.Context, userID, provider string, token *oauth2.Token) error
}

type Storage interface {
	TokenReader
	TokenStorer
}
