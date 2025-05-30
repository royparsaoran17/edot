package presentations

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Login) Validate() error {
	return validation.Errors{
		"username": validation.Validate(&r.Username, validation.Required),
		"password": validation.Validate(&r.Password, validation.Required),
	}.Filter()
}

type Verify struct {
	Token string `json:"token"`
}

func (r *Verify) Validate() error {
	return validation.Errors{
		"token": validation.Validate(&r.Token, validation.Required),
	}.Filter()
}
