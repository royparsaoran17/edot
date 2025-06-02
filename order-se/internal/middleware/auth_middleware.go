package middleware

import (
	"net/http"
	"order-se/internal/appctx"
)

func Authorize() MiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, conf *appctx.Config) error {
		return nil
	}
}
