package router

import (
	usersvc "auth-se/internal/service/user"
	"auth-se/internal/ucase/user"
	"net/http"

	"auth-se/internal/handler"
)

func (rtr *router) mountUser(userSvc usersvc.User) {
	rtr.router.HandleFunc("/internal/v1/users", rtr.handle(
		handler.HttpRequest,
		user.NewUserGetAll(userSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/users", rtr.handle(
		handler.HttpRequest,
		user.NewUserCreate(userSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/users/{user_id}", rtr.handle(
		handler.HttpRequest,
		user.NewUserGetByID(userSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/users/{user_id}", rtr.handle(
		handler.HttpRequest,
		user.NewUserUpdate(userSvc),
	)).Methods(http.MethodPut)

}
