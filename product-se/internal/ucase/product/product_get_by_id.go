package product

import (
	"net/http"

	"product-se/pkg/tracer"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/service/product"
	"product-se/internal/ucase/contract"
)

type productGetByID struct {
	service product.Product
}

func (r productGetByID) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.product_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("productGetByID")

	productID := mux.Vars(data.Request)["product_id"]
	if _, err := uuid.Parse(productID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	result, err := r.service.GetProductByID(ctx, productID)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrProductNotFound:
			return *responder.
				WithCode(http.StatusNotFound).
				WithMessage(causer.Error())

		default:
			tracer.SpanError(ctx, err)
			return *responder.
				WithCode(http.StatusInternalServerError).
				WithMessage(http.StatusText(http.StatusInternalServerError))
		}

	}

	return *responder.
		WithData(result).
		WithCode(http.StatusOK).
		WithMessage("product fetched")
}

func NewProductGetByID(service product.Product) contract.UseCase {
	return &productGetByID{service: service}
}
