package product

import (
	"net/http"
	"product-se/internal/common"
	"product-se/pkg/tracer"

	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/service/product"
	"product-se/internal/ucase/contract"
)

type productGetAll struct {
	service product.Product
}

func (r productGetAll) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.product_get_all")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("productGetAll")

	metaData := common.MetadataFromURL(data.Request.URL.Query())

	metaDateRange, err := common.DateRangeFromURL(data.Request.URL.Query(), "created_at", "created_from", "created_until")
	if err != nil {
		return *responder.WithCode(http.StatusBadRequest).WithMessage(err.Error())
	}

	metaData.DateRange = metaDateRange

	products, err := r.service.GetAllProduct(ctx, &metaData)
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
		WithMessage("get all products success").
		WithData(products)
}

func NewProductGetAll(service product.Product) contract.UseCase {
	return &productGetAll{service: service}
}
