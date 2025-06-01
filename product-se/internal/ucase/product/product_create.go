package product

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"product-se/internal/consts"

	"product-se/pkg/tracer"

	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/presentations"
	"product-se/internal/service/product"
	"product-se/internal/ucase/contract"
)

type productCreate struct {
	service product.Product
}

func (r productCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.product_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("productCreate")

	var input presentations.ProductCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	_, err := r.service.CreateProduct(ctx, input)
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
		WithMessage("product created")
}

func NewProductCreate(service product.Product) contract.UseCase {
	return &productCreate{service: service}
}
