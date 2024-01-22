package gateway

import (
	"go.datalift.io/datalift/server/endpoint"
	"go.datalift.io/datalift/server/middleware"
	"go.datalift.io/datalift/server/service"
)

var Middleware = middleware.Factory{}

var Endpoints = endpoint.Factory{}

var Services = service.Factory{}

var CoreComponentFactory = &ComponentFactory{
	Services:   Services,
	Middleware: Middleware,
	Endpoints:  Endpoints,
}
