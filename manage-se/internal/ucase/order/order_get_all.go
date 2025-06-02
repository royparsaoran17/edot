package order

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"manage-se/internal/common"
	"manage-se/internal/entity"
	"manage-se/internal/provider/providererrors"
	"manage-se/pkg/logger"
	"manage-se/pkg/tracer"
	"net/http"

	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/consts"
	"manage-se/internal/service/order"
	"manage-se/internal/ucase/contract"
)

type orderGetAll struct {
	service order.Order
}

func (r orderGetAll) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.order_get_all")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("orderGetAll")

	metaData := common.MetadataFromURL(data.Request.URL.Query())

	metaDateRange, err := common.DateRangeFromURL(data.Request.URL.Query(), "created_at", "created_from", "created_until")
	if err != nil {
		return *responder.WithCode(http.StatusBadRequest).WithMessage(err.Error())
	}

	user, ok := ctx.Value(consts.CtxUserAuth).(entity.UserContext)
	if !ok {
		logger.Error(errors.New("user data not exist in context"))
		return *responder.WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
	}

	metaData.DateRange = metaDateRange

	orders, err := r.service.GetAllOrder(ctx, user.ID, &metaData)
	if err != nil {
		errCause := errors.Cause(err)
		switch errCause {
		default:
			switch caorder := errCause.(type) {
			case consts.Error:
				return *responder.WithContext(ctx).WithCode(http.StatusBadRequest).WithMessage(errCause.Error())

			case providererrors.Error:
				return *responder.WithContext(ctx).WithCode(caorder.Code).WithError(caorder.Errors).WithMessage(caorder.Message)

			case validation.Errors:
				return *responder.
					WithContext(ctx).
					WithCode(http.StatusUnprocessableEntity).
					WithMessage("Validation Error(s)").
					WithError(errCause)

			default:
				return *responder.WithContext(ctx).WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
			}
		}
	}

	return *responder.
		WithCode(http.StatusOK).
		WithMessage("get all orders success").
		WithMeta(metaData).
		WithData(orders)
}

func NewOrderGetAll(service order.Order) contract.UseCase {
	return &orderGetAll{service: service}
}
