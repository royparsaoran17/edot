package warehouse

import (
	"net/http"
	"product-se/internal/common"
	"product-se/pkg/tracer"

	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/service/warehouse"
	"product-se/internal/ucase/contract"
)

type warehouseGetAll struct {
	service warehouse.Warehouse
}

func (r warehouseGetAll) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_get_all")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseGetAll")

	metaData := common.MetadataFromURL(data.Request.URL.Query())

	metaDateRange, err := common.DateRangeFromURL(data.Request.URL.Query(), "created_at", "created_from", "created_until")
	if err != nil {
		return *responder.WithCode(http.StatusBadRequest).WithMessage(err.Error())
	}

	metaData.DateRange = metaDateRange

	warehouses, err := r.service.GetAllWarehouse(ctx, &metaData)
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
		WithMessage("get all warehouses success").
		WithData(warehouses)
}

func NewWarehouseGetAll(service warehouse.Warehouse) contract.UseCase {
	return &warehouseGetAll{service: service}
}
