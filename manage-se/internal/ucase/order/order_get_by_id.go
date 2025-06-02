package order

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/consts"
	"manage-se/internal/service/order"
	"manage-se/internal/ucase/contract"
	"manage-se/pkg/tracer"
)

type orderGetByID struct {
	service order.Order
}

func (r orderGetByID) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("orderGetByID")

	orderID := mux.Vars(data.Request)["order_id"]
	if _, err := uuid.Parse(orderID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	result, err := r.service.GetOrderByID(ctx, orderID)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrOrderNotFound:
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
		WithMessage("order fetched")
}

func NewOrderGetByID(service order.Order) contract.UseCase {
	return &orderGetByID{service: service}
}
