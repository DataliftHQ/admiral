package service

import (
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
)

type Service interface{}

type Factory map[string]func(*anypb.Any, *zap.Logger, tally.Scope) (Service, error)

var Registry = map[string]Service{}
