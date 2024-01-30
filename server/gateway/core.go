package gateway

import (
	"go.datalift.io/admiral/server/endpoint"
	healthcheckendp "go.datalift.io/admiral/server/endpoint/healthcheck"
	"go.datalift.io/admiral/server/middleware"
	"go.datalift.io/admiral/server/middleware/stats"
	"go.datalift.io/admiral/server/middleware/validate"
	"go.datalift.io/admiral/server/service"
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
