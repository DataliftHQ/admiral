package gateway

import (
	"go.datalift.io/admiral/server/endpoint"
	authnendp "go.datalift.io/admiral/server/endpoint/authn"
	fooendp "go.datalift.io/admiral/server/endpoint/foo"
	healthcheckendp "go.datalift.io/admiral/server/endpoint/healthcheck"
	metaendp "go.datalift.io/admiral/server/endpoint/meta"
	settingsendp "go.datalift.io/admiral/server/endpoint/settings"
	usersendp "go.datalift.io/admiral/server/endpoint/users"
	"go.datalift.io/admiral/server/middleware"
	"go.datalift.io/admiral/server/middleware/authn"
	"go.datalift.io/admiral/server/middleware/stats"
	"go.datalift.io/admiral/server/middleware/validate"
	"go.datalift.io/admiral/server/service"
	authnservice "go.datalift.io/admiral/server/service/authn"
	"go.datalift.io/admiral/server/service/db/postgres"
	"go.datalift.io/admiral/server/service/temporal"
)

var Middleware = middleware.Factory{
	stats.Name:    stats.New,
	validate.Name: validate.New,
	authn.Name:    authn.New,
}

var Endpoints = endpoint.Factory{
	healthcheckendp.Name: healthcheckendp.New,
	settingsendp.Name:    settingsendp.New,
	fooendp.Name:         fooendp.New,
	metaendp.Name:        metaendp.New,
	usersendp.Name:       usersendp.New,
	authnendp.Name:       authnendp.New,
}

var Services = service.Factory{
	postgres.Name:     postgres.New,
	temporal.Name:     temporal.New,
	authnservice.Name: authnservice.New,
}

var CoreComponentFactory = &ComponentFactory{
	Services:   Services,
	Middleware: Middleware,
	Endpoints:  Endpoints,
}
