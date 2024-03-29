package accesslog

import (
	"context"
	"fmt"

	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	accesslogv1 "go.datalift.io/admiral/server/config/middleware/accesslog/v1"
	"go.datalift.io/admiral/server/gateway/log"
	"go.datalift.io/admiral/server/gateway/meta"
	"go.datalift.io/admiral/server/middleware"
)

const Name = "admiral.middleware.accesslog"

func New(config *accesslogv1.Config, logger *zap.Logger, scope tally.Scope) (middleware.Middleware, error) {
	var statusCodes []codes.Code
	// if no filter is provided default to logging all status codes
	if config != nil {
		for _, filter := range config.StatusCodeFilters {
			switch t := filter.GetFilterType().(type) {
			case *accesslogv1.Config_StatusCodeFilter_Equals:
				statusCode := filter.GetEquals()
				statusCodes = append(statusCodes, codes.Code(statusCode))
			default:
				return nil, fmt.Errorf("status code filter `%T` not supported", t)
			}
		}
	}
	return &mid{
		logger:      logger,
		statusCodes: statusCodes,
	}, nil
}

type mid struct {
	logger *zap.Logger
	// TODO(perf): improve lookup efficiency using a lookup table
	statusCodes []codes.Code
}

func (m *mid) UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		service, method, ok := middleware.SplitFullMethod(info.FullMethod)
		if !ok {
			m.logger.Warn("could not parse gRPC method", zap.String("fullMethod", info.FullMethod))
		}
		resp, err := handler(ctx, req)
		s := status.Convert(err)
		if s == nil {
			s = status.New(codes.OK, "")
		}
		code := s.Code()
		// api logger context fields
		fields := []zap.Field{
			zap.String("service", service),
			zap.String("method", method),
			zap.Int("statusCode", int(code)),
			zap.String("status", code.String()),
		}

		if m.validStatusCode(code) {
			// if err is returned from handler, log error details only
			// as response body will be nil
			if err != nil {
				reqBody, err := meta.APIBody(req)
				if err != nil {
					return nil, err
				}
				fields = append(fields, log.ProtoField("requestBody", reqBody))
				fields = append(fields, zap.String("error", s.Message()))
				m.logger.Error("gRPC", fields...)
			} else {
				m.logger.Info("gRPC", fields...)
			}
		}
		return resp, err
	}
}

func (m *mid) validStatusCode(c codes.Code) bool {
	// If no filter is provided all status codes are valid
	if len(m.statusCodes) == 0 {
		return true
	}
	for _, code := range m.statusCodes {
		if c == code {
			return true
		}
	}
	return false
}
