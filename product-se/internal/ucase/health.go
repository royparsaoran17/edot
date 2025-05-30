package ucase

import (
	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/ucase/contract"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("ok")
}
