package user

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"manage-se/internal/common"
	"manage-se/internal/provider/providererrors"
	"manage-se/pkg/tracer"
	"net/http"

	"github.com/pkg/errors"
	"manage-se/internal/appctx"
	"manage-se/internal/consts"
	"manage-se/internal/service/user"
	"manage-se/internal/ucase/contract"
)

type userGetAll struct {
	service user.User
}

func (r userGetAll) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.user_get_all")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("userGetAll")

	metaData := common.MetadataFromURL(data.Request.URL.Query())

	metaDateRange, err := common.DateRangeFromURL(data.Request.URL.Query(), "created_at", "created_from", "created_until")
	if err != nil {
		return *responder.WithCode(http.StatusBadRequest).WithMessage(err.Error())
	}

	metaData.DateRange = metaDateRange

	users, err := r.service.GetAllUser(ctx, &metaData)
	if err != nil {
		errCause := errors.Cause(err)
		switch errCause {
		default:
			switch causer := errCause.(type) {
			case consts.Error:
				return *responder.WithContext(ctx).WithCode(http.StatusBadRequest).WithMessage(errCause.Error())

			case providererrors.Error:
				return *responder.WithContext(ctx).WithCode(causer.Code).WithError(causer.Errors).WithMessage(causer.Message)

			case validation.Errors:
				return *responder.
					WithContext(ctx).
					WithCode(http.StatusUnprocessableEntity).
					WithMessage("Validation Error(s)").
					WithError(errCause)

			default:
				return *responder.WithContext(ctx).WithCode(http.StatusInternalServerError).WithMessage(http.StatusText(http.StatusInternalServerError))
			}
		}
	}

	return *responder.
		WithCode(http.StatusOK).
		WithMessage("get all users success").
		WithMeta(metaData).
		WithData(users)
}

func NewUserGetAll(service user.User) contract.UseCase {
	return &userGetAll{service: service}
}
