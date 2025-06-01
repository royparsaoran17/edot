package product

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/presentations"
	"product-se/internal/service/product"
	"product-se/internal/ucase/contract"
	"product-se/pkg/logger"
	"product-se/pkg/tracer"
)

type productUpdate struct {
	service product.Product
}

func (r productUpdate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.product_update")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("productUpdate")
	var input presentations.ProductUpdate

	if err := data.Cast(&input); err != nil {
		logger.Warn(fmt.Sprintf("error cast to productUpdate presentation %+v", err))
		tracer.SpanError(ctx, err)
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	productID := mux.Vars(data.Request)["product_id"]
	if _, err := uuid.Parse(productID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	err := r.service.UpdateProductByID(ctx, productID, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrProductNotFound:
			return *responder.
				WithCode(http.StatusNotFound).
				WithMessage(causer.Error())

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
		WithCode(http.StatusOK).
		WithMessage("product updated")
}

func NewProductUpdate(service product.Product) contract.UseCase {
	return &productUpdate{service: service}
}
