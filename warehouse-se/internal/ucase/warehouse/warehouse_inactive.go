package warehouse

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"warehouse-se/internal/appctx"
	"warehouse-se/internal/consts"
	"warehouse-se/internal/service/warehouse"
	"warehouse-se/internal/ucase/contract"
	"warehouse-se/pkg/tracer"
)

type warehouseInactive struct {
	service warehouse.Warehouse
}

func (r warehouseInactive) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_inactive")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseInactive")
	warehouseID := mux.Vars(data.Request)["warehouse_id"]
	if _, err := uuid.Parse(warehouseID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	err := r.service.InactiveWarehouseByID(ctx, warehouseID)
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
		WithMessage("warehouse inactive")
}

func NewWarehouseInactive(service warehouse.Warehouse) contract.UseCase {
	return &warehouseInactive{service: service}
}
