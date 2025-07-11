// Package handler
package handler

import (
	"context"
	"net/http"

	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/ucase/contract"
)

// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response {
	ctx := context.WithValue(request.Context(), consts.CtxLang, request.Header.Get(consts.HeaderLanguageKey))

	req := request.WithContext(ctx)

	data := &appctx.Data{
		Request:     req,
		Config:      conf,
		ServiceType: consts.ServiceTypeHTTP,
	}

	return svc.Serve(data)
}
