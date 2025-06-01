package warehouse

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"product-se/internal/consts"

	"product-se/pkg/tracer"

	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/presentations"
	"product-se/internal/service/warehouse"
	"product-se/internal/ucase/contract"
)

type warehouseMoveStock struct {
	service warehouse.Warehouse
}

func (r warehouseMoveStock) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_move_stock")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseMoveStock")

	warehouseID := mux.Vars(data.Request)["warehouse_id"]
	if _, err := uuid.Parse(warehouseID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	var input presentations.WarehouseCreateMoveStock
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	input.FromWarehouseID = warehouseID
	err := r.service.MoveWarehouseStock(ctx, input)
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
		WithMessage("warehouse moveed successfully")
}

func NewWarehouseMoveStock(service warehouse.Warehouse) contract.UseCase {
	return &warehouseMoveStock{service: service}
}
