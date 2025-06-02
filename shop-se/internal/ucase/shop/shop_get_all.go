package shop

import (
	"net/http"
	"shop-se/internal/common"
	"shop-se/pkg/tracer"

	"github.com/pkg/errors"
	"shop-se/internal/appctx"
	"shop-se/internal/consts"
	"shop-se/internal/service/shop"
	"shop-se/internal/ucase/contract"
)

type shopGetAll struct {
	service shop.Shop
}

func (r shopGetAll) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.shop_get_all")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("shopGetAll")

	metaData := common.MetadataFromURL(data.Request.URL.Query())

	metaDateRange, err := common.DateRangeFromURL(data.Request.URL.Query(), "created_at", "created_from", "created_until")
	if err != nil {
		return *responder.WithCode(http.StatusBadRequest).WithMessage(err.Error())
	}

	metaData.DateRange = metaDateRange

	shops, err := r.service.GetAllShop(ctx, &metaData)
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
		WithMessage("get all shops success").
		WithData(shops)
}

func NewShopGetAll(service shop.Shop) contract.UseCase {
	return &shopGetAll{service: service}
}
