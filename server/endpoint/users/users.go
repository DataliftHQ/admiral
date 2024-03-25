package users

import (
	"context"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	usersv1 "go.datalift.io/admiral/common/api/users/v1"
	"go.datalift.io/admiral/server/endpoint"
)

const (
	Name = "admiral.endpoint.users"
)

type endp struct {
	logger *zap.Logger
	scope  tally.Scope
}

func New(_ *anypb.Any, log *zap.Logger, scope tally.Scope) (endpoint.Endpoint, error) {
	return &endp{
		logger: log,
		scope:  scope,
	}, nil
}

func (e *endp) Register(r endpoint.Registrar) error {
	usersv1.RegisterUsersAPIServer(r.GRPCServer(), e)
	return r.RegisterJSONGateway(usersv1.RegisterUsersAPIHandler)
}

func (e *endp) GetMe(ctx context.Context, req *usersv1.GetMeRequest) (*usersv1.GetMeResponse, error) {
	return &usersv1.GetMeResponse{
		Id:         "1",
		GivenName:  "Martin",
		FamilyName: "Berwanger",
		Email:      "mberwanger@protonmail.com",
	}, nil
}

func (e *endp) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	return &usersv1.GetUserResponse{
		Id:         "1",
		GivenName:  "Martin",
		FamilyName: "Berwanger",
		Email:      "mberwanger@protonmail.com",
	}, nil
}
