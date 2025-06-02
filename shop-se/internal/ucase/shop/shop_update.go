package shop

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"shop-se/internal/appctx"
	"shop-se/internal/consts"
	"shop-se/internal/presentations"
	"shop-se/internal/service/shop"
	"shop-se/internal/ucase/contract"
	"shop-se/pkg/logger"
	"shop-se/pkg/tracer"
)

type shopUpdate struct {
	service shop.Shop
}

func (r shopUpdate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.shop_update")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("shopUpdate")
	var input presentations.ShopUpdate

	if err := data.Cast(&input); err != nil {
		logger.Warn(fmt.Sprintf("error cast to shopUpdate presentation %+v", err))
		tracer.SpanError(ctx, err)
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	shopID := mux.Vars(data.Request)["shop_id"]
	if _, err := uuid.Parse(shopID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	err := r.service.UpdateShopByID(ctx, shopID, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrShopNotFound:
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
		WithMessage("shop updated")
}

func NewShopUpdate(service shop.Shop) contract.UseCase {
	return &shopUpdate{service: service}
}
