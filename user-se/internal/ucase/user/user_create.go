package user

import (
	"auth-se/internal/consts"
	"auth-se/internal/service/user"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"net/http"

	"auth-se/pkg/tracer"

	"auth-se/internal/appctx"
	"auth-se/internal/presentations"
	"auth-se/internal/ucase/contract"
	"github.com/pkg/errors"
)

type userCreate struct {
	service user.User
}

func (c userCreate) Serve(data *appctx.Data) appctx.Response {
	ctx := tracer.SpanStart(data.Request.Context(), "usecase.user_create")
	defer tracer.SpanFinish(ctx)

	responder := appctx.NewResponse().WithState("userCreate")

	var input presentations.UserCreate
	if err := data.Cast(&input); err != nil {
		return *responder.WithCode(http.StatusBadRequest).
			WithError(err.Error()).
			WithMessage(http.StatusText(http.StatusBadRequest))
	}

	users, err := c.service.CreateUser(ctx, input)
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
		WithData(users).
		WithMessage("user created")
}

func NewUserCreate(service user.User) contract.UseCase {
	return &userCreate{service: service}
}
