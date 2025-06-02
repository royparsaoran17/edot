package warehouse

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"warehouse-se/internal/consts"

	"warehouse-se/pkg/tracer"

	"github.com/pkg/errors"
	"warehouse-se/internal/appctx"
	"warehouse-se/internal/presentations"
	"warehouse-se/internal/service/warehouse"
	"warehouse-se/internal/ucase/contract"
)

type warehouseAddStock struct {
	service warehouse.Warehouse
}

func (r warehouseAddStock) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_add_stock")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseAddStock")

	warehouseID := mux.Vars(data.Request)["warehouse_id"]
	if _, err := uuid.Parse(warehouseID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	var input presentations.WarehouseStock
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	input.ID = warehouseID
	_, err := r.service.AddWarehouseStock(ctx, input)
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
		WithCode(http.StatusOK).
		WithMessage("warehouse added successfully")
}

func NewWarehouseAddStock(service warehouse.Warehouse) contract.UseCase {
	return &warehouseAddStock{service: service}
}
