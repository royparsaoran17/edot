package order

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"order-se/internal/appctx"
	"order-se/internal/consts"
	"order-se/internal/presentations"
	"order-se/internal/service/order"
	"order-se/internal/ucase/contract"
	"order-se/pkg/logger"
	"order-se/pkg/tracer"
)

type orderPayment struct {
	service order.Order
}

func (r orderPayment) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_payment")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("orderPayment")
	var input presentations.OrderPayment

	if err := data.Cast(&input); err != nil {
		logger.Warn(fmt.Sprintf("error cast to orderPayment presentation %+v", err))
		tracer.SpanError(ctx, err)
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	orderID := mux.Vars(data.Request)["order_id"]
	if _, err := uuid.Parse(orderID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	input.OrderID = orderID
	_, err := r.service.CreateOrderPayment(ctx, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrOrderNotFound:
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
		WithMessage("order payment")
}

func NewOrderPayment(service order.Order) contract.UseCase {
	return &orderPayment{service: service}
}
