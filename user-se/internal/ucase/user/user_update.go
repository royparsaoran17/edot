package user

import (
	"auth-se/internal/service/user"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"auth-se/internal/appctx"
	"auth-se/internal/consts"
	"auth-se/internal/presentations"
	"auth-se/internal/ucase/contract"
	"auth-se/pkg/logger"
	"auth-se/pkg/tracer"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type userUpdate struct {
	service user.User
}

func (c userUpdate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.user_update")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("userUpdate")
	var input presentations.UserUpdate

	if err := data.Cast(&input); err != nil {
		logger.Warn(fmt.Sprintf("error cast to userUpdate presentation %+v", err))
		tracer.SpanError(ctx, err)
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	userID := mux.Vars(data.Request)["user_id"]
	if _, err := uuid.Parse(userID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	err := c.service.UpdateUserByID(ctx, userID, input)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrUserNotFound:
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
		WithMessage("user updated")
}

func NewUserUpdate(service user.User) contract.UseCase {
	return &userUpdate{service: service}
}
