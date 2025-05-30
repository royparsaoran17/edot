// Package router
package server

import (
	"context"
	"net/http"

	"product-se/internal/appctx"
	ucase "product-se/internal/ucase/contract"
)

// httpHandlerFunc abstraction for http handler
type httpHandlerFunc func(request *http.Request, svc ucase.UseCase, conf *appctx.Config) appctx.Response

// Server contract
type Server interface {
	Run(context.Context) error
	Done()
	Config() *appctx.Config
}
