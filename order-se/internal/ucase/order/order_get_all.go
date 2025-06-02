package order

import (
	"net/http"
	"order-se/internal/common"
	"order-se/pkg/tracer"

	"github.com/pkg/errors"
	"order-se/internal/appctx"
	"order-se/internal/consts"
	"order-se/internal/service/order"
	"order-se/internal/ucase/contract"
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

	metaData.DateRange = metaDateRange

	userID := data.Request.URL.Query().Get("user_id")
	orders, err := r.service.GetAllOrder(ctx, userID, &metaData)
	if err != nil {

		switch causer := errors.Cause(err); causer {
		case common.ErrInvalidMetadata:
			return *responder.
				WithCode(http.StatusBadRequest).
				WithMessage(err.Error())

		default:
			switch causer.(type) {
			case consts.Error:
				return *responder.
					WithCode(http.StatusBadRequest).
					WithMessage(causer.Error())

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
		WithMessage("get all orders success").
		WithData(orders)
}

func NewOrderGetAll(service order.Order) contract.UseCase {
	return &orderGetAll{service: service}
}
