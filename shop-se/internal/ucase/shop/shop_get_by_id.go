package shop

import (
	"net/http"

	"shop-se/pkg/tracer"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"shop-se/internal/appctx"
	"shop-se/internal/consts"
	"shop-se/internal/service/shop"
	"shop-se/internal/ucase/contract"
)

type shopGetByID struct {
	service shop.Shop
}

func (r shopGetByID) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.shop_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("shopGetByID")

	shopID := mux.Vars(data.Request)["shop_id"]
	if _, err := uuid.Parse(shopID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	result, err := r.service.GetShopByID(ctx, shopID)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrShopNotFound:
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
		WithMessage("shop fetched")
}

func NewShopGetByID(service shop.Shop) contract.UseCase {
	return &shopGetByID{service: service}
}
