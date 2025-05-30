package user

import (
	"auth-se/internal/common"
	"auth-se/internal/service/user"
	"auth-se/pkg/tracer"
	"net/http"

	"auth-se/internal/appctx"
	"auth-se/internal/consts"
	"auth-se/internal/ucase/contract"
	"github.com/pkg/errors"
)

type userGetAll struct {
	service user.User
}

func (c userGetAll) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.user_get_all")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("userGetAll")

	metaData := common.MetadataFromURL(data.Request.URL.Query())

	metaDateRange, err := common.DateRangeFromURL(data.Request.URL.Query(), "created_at", "created_from", "created_until")
	if err != nil {
		return *responder.WithCode(http.StatusBadRequest).WithMessage(err.Error())
	}

	metaData.DateRange = metaDateRange

	users, err := c.service.GetAllUser(ctx, &metaData)
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
		WithMeta(metaData).
		WithMessage("get all users success").
		WithData(users)
}

func NewUserGetAll(service user.User) contract.UseCase {
	return &userGetAll{service: service}
}
