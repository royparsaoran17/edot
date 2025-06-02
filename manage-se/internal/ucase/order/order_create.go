package order

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/consts"
	"manage-se/internal/entity"
	"manage-se/internal/presentations"
	"manage-se/internal/service/order"
	"manage-se/internal/ucase/contract"
	"manage-se/pkg/logger"

	"manage-se/pkg/tracer"
	"net/http"
)

type orderCreate struct {
	service order.Order
}

func (r orderCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("orderCreate")

	var input presentations.OrderCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	user, ok := ctx.Value(consts.CtxUserAuth).(entity.UserContext)
	if !ok {
		logger.Error(errors.New("user data not exist in context"))
		return *responder.WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
	}

	input.UserID = user.ID
	_, err := r.service.CreateOrder(ctx, input)
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
		WithMessage("order created")
}

func NewOrderCreate(service order.Order) contract.UseCase {
	return &orderCreate{service: service}
}
