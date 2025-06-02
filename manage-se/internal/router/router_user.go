package router

import (
	"manage-se/internal/middleware"
	authsvc "manage-se/internal/service/auth"
	"manage-se/internal/ucase/user"
	"net/http"

	"manage-se/internal/handler"
	usersvc "manage-se/internal/service/user"
)

func (rtr *router) mountUser(userSvc usersvc.User, authSvc authsvc.Auth) {
	rtr.router.HandleFunc("/external/v1/users", rtr.handle(
		handler.HttpRequest,
		user.NewUserGetAll(userSvc),
		middleware.Authorize(authSvc, "Super Admin"),
	)).Methods(http.MethodGet)

}
