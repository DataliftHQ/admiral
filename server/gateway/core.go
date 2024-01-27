package gateway

import (
	"go.datalift.io/datalift/server/endpoint"
	healthcheckendp "go.datalift.io/datalift/server/endpoint/healthcheck"
	"go.datalift.io/datalift/server/middleware"
	"go.datalift.io/datalift/server/middleware/stats"
	"go.datalift.io/datalift/server/middleware/validate"
	"go.datalift.io/datalift/server/service"
)

var Middleware = middleware.Factory{
	stats.Name:    stats.New,
	validate.Name: validate.New,
}

var Endpoints = endpoint.Factory{
	healthcheckendp.Name: healthcheckendp.New,
}

var Services = service.Factory{}

var CoreComponentFactory = &ComponentFactory{
	Services:   Services,
	Middleware: Middleware,
	Endpoints:  Endpoints,
}
