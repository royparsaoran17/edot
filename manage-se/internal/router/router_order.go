package router

import (
	"manage-se/internal/consts"
	"manage-se/internal/middleware"
	authsvc "manage-se/internal/service/auth"
	"manage-se/internal/ucase/order"
	"net/http"

	"manage-se/internal/handler"
	ordersvc "manage-se/internal/service/order"
)

func (rtr *router) mountOrder(orderSvc ordersvc.Order, authSvc authsvc.Auth) {
	rtr.router.HandleFunc("/external/v1/orders", rtr.handle(
		handler.HttpRequest,
		order.NewOrderGetAll(orderSvc),
		middleware.Authorize(authSvc, consts.AllRoles),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/external/v1/orders", rtr.handle(
		handler.HttpRequest,
		order.NewOrderCreate(orderSvc),
		middleware.Authorize(authSvc, consts.AllRoles),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/external/v1/orders/{order_id}", rtr.handle(
		handler.HttpRequest,
		order.NewOrderGetByID(orderSvc),
		middleware.Authorize(authSvc, consts.AllRoles),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/external/v1/orders/{order_id}/payment", rtr.handle(
		handler.HttpRequest,
		order.NewOrderPayment(orderSvc),
		middleware.Authorize(authSvc, consts.AllRoles),
	)).Methods(http.MethodPost)

}
