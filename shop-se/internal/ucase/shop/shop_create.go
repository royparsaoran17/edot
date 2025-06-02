package shop

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"shop-se/internal/consts"

	"shop-se/pkg/tracer"

	"github.com/pkg/errors"
	"shop-se/internal/appctx"
	"shop-se/internal/presentations"
	"shop-se/internal/service/shop"
	"shop-se/internal/ucase/contract"
)

type shopCreate struct {
	service shop.Shop
}

func (r shopCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.shop_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("shopCreate")

	var input presentations.ShopCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	_, err := r.service.CreateShop(ctx, input)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		default:
			switch cause := causer.(type) {
			case consts.Error:
				return *responder.
					WithCode(http.StatusBadRequest).
					WithMessage(cause.Error())

			case validation.Errors:
				return *responder.
					WithCode(http.StatusUnprocessableEntity).
					WithError(cause).
					WithMessage("Validation error(s)")

			default:
				tracer.SpanError(ctx, err)
				return *responder.
					WithCode(http.StatusInternalServerError).
					WithMessage(http.StatusText(http.StatusInternalServerError))
			}
		}
	}

	return *responder.
		WithCode(http.StatusCreated).
		WithMessage("shop created")
}

func NewShopCreate(service shop.Shop) contract.UseCase {
	return &shopCreate{service: service}
}
