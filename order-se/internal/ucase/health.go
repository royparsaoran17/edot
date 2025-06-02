package ucase

import (
	"order-se/internal/appctx"
	"order-se/internal/consts"
	"order-se/internal/ucase/contract"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return *appctx.NewResponse().WithCode(consts.CodeSuccess).WithMessage("ok")
}
