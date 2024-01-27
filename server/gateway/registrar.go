package gateway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"go.datalift.io/datalift/server/endpoint"
)

func newRegistrar(ctx context.Context, m *runtime.ServeMux, s *grpc.Server, c *grpc.ClientConn) endpoint.Registrar {
	return &registrar{
		ctx: ctx,
		s:   s,
		m:   m,
		c:   c,
	}
}

type registrar struct {
	ctx context.Context

	s *grpc.Server
	c *grpc.ClientConn
	m *runtime.ServeMux
}

func (r *registrar) GRPCServer() *grpc.Server {
	return r.s
}

func (r *registrar) RegisterJSONGateway(f endpoint.GatewayRegisterAPIHandlerFunc) error {
	return f(r.ctx, r.m, r.c)
}
