package warehouse

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"
	"warehouse-se/internal/consts"

	"warehouse-se/pkg/tracer"

	"github.com/pkg/errors"
	"warehouse-se/internal/appctx"
	"warehouse-se/internal/presentations"
	"warehouse-se/internal/service/warehouse"
	"warehouse-se/internal/ucase/contract"
)

type warehouseCreate struct {
	service warehouse.Warehouse
}

func (r warehouseCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.warehouse_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("warehouseCreate")

	var input presentations.WarehouseCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	_, err := r.service.CreateWarehouse(ctx, input)
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
		WithCode(http.StatusCreated).
		WithMessage("warehouse created")
}

func NewWarehouseCreate(service warehouse.Warehouse) contract.UseCase {
	return &warehouseCreate{service: service}
}
