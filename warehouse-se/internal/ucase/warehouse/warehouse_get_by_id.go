package warehouse

import (
	"net/http"

	"warehouse-se/pkg/tracer"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"warehouse-se/internal/appctx"
	"warehouse-se/internal/consts"
	"warehouse-se/internal/service/warehouse"
	"warehouse-se/internal/ucase/contract"
)

type warehouseGetByID struct {
	service warehouse.Warehouse
}

func (r warehouseGetByID) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseGetByID")

	warehouseID := mux.Vars(data.Request)["warehouse_id"]
	if _, err := uuid.Parse(warehouseID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	result, err := r.service.GetWarehouseByID(ctx, warehouseID)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrWarehouseNotFound:
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
		WithMessage("warehouse fetched")
}

func NewWarehouseGetByID(service warehouse.Warehouse) contract.UseCase {
	return &warehouseGetByID{service: service}
}
