package meta

import (
	"context"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"

	metav1 "go.datalift.io/admiral/common/api/meta/v1"
	"go.datalift.io/admiral/server/endpoint"
)

const (
	Name = "admiral.endpoint.meta"
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
	metav1.RegisterMetaAPIServer(r.GRPCServer(), e)
	return r.RegisterJSONGateway(metav1.RegisterMetaAPIHandler)
}

func (e *endp) GetMeta(ctx context.Context, req *metav1.GetMetaRequest) (*metav1.GetMetaResponse, error) {
	return &metav1.GetMetaResponse{}, nil
}
