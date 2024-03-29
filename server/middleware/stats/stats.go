package stats

import (
	"context"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"go.datalift.io/admiral/server/middleware"
)

const Name = "admiral.middleware.stats"

func New(cfg *anypb.Any, logger *zap.Logger, scope tally.Scope) (middleware.Middleware, error) {
	return &mid{
		logger: logger,
		scope:  scope.SubScope("module"),
	}, nil
}

type mid struct {
	logger *zap.Logger
	scope  tally.Scope
}

func (m *mid) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		service, method, ok := middleware.SplitFullMethod(info.FullMethod)
		if !ok {
			m.logger.Warn("could not parse gRPC method", zap.String("fullMethod", info.FullMethod))
		}

		grpcScope := m.scope.Tagged(map[string]string{
			"grpc_service": service,
			"grpc_method":  method,
		})

		t := grpcScope.Timer("rpc_latency").Start()
		resp, err := handler(ctx, req)
		t.Stop()

		grpcScope.Tagged(map[string]string{
			"grpc_status": status.Convert(err).Code().String(),
		}).Counter("rpc_total").Inc(1)

		return resp, err
	}
}
