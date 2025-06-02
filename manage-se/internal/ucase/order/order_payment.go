package order

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"manage-se/internal/presentations"
	"manage-se/pkg/logger"
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

type orderPayment struct {
	service order.Order
}

func (r orderPayment) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_payment")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("orderPayment")
	var input presentations.PaymentCreate

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
	_, err := r.service.CreateOrderPayment(ctx, orderID, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrDataNotFound:
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
