package validate

import (
	validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"

	"go.datalift.io/datalift/server/middleware"
)

const Name = "clutch.middleware.validate"

func New(cfg *anypb.Any, logger *zap.Logger, scope tally.Scope) (middleware.Middleware, error) {
	return &mid{}, nil
}

type mid struct{}

func (m *mid) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return validator.UnaryServerInterceptor()
}
