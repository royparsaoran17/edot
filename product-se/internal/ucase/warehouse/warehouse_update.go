package warehouse

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"product-se/internal/appctx"
	"product-se/internal/consts"
	"product-se/internal/presentations"
	"product-se/internal/service/warehouse"
	"product-se/internal/ucase/contract"
	"product-se/pkg/logger"
	"product-se/pkg/tracer"
)

type warehouseUpdate struct {
	service warehouse.Warehouse
}

func (r warehouseUpdate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_update")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseUpdate")
	var input presentations.WarehouseUpdate

	if err := data.Cast(&input); err != nil {
		logger.Warn(fmt.Sprintf("error cast to warehouseUpdate presentation %+v", err))
		tracer.SpanError(ctx, err)
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	warehouseID := mux.Vars(data.Request)["warehouse_id"]
	if _, err := uuid.Parse(warehouseID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	err := r.service.UpdateWarehouseByID(ctx, warehouseID, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrWarehouseNotFound:
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
		WithMessage("warehouse updated")
}

func NewWarehouseUpdate(service warehouse.Warehouse) contract.UseCase {
	return &warehouseUpdate{service: service}
}
