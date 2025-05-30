package user

import (
	"auth-se/internal/service/user"
	"net/http"

	"auth-se/pkg/tracer"

	"auth-se/internal/appctx"
	"auth-se/internal/consts"
	"auth-se/internal/ucase/contract"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type userGetByID struct {
	service user.User
}

func (c userGetByID) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.user_get_by_id")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("userGetByID")

	userID := mux.Vars(data.Request)["user_id"]
	if _, err := uuid.Parse(userID); err != nil {
		return *responder.
			WithCode(http.StatusBadRequest).
			WithMessage(consts.ErrInvalidUUID.Error())
	}

	result, err := c.service.GetUserByID(ctx, userID)
	if err != nil {
		switch causer := errors.Cause(err); causer {
		case consts.ErrUserNotFound:
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
		WithMessage("user fetched")
}

func NewUserGetByID(service user.User) contract.UseCase {
	return &userGetByID{service: service}
}
