package router

import (
	"net/http"
	"order-se/internal/handler"
	"order-se/internal/ucase/order"

	ordersvc "order-se/internal/service/order"
)

func (rtr *router) mountOrders(orderSvc ordersvc.Order) {
	rtr.router.HandleFunc("/internal/v1/orders", rtr.handle(
		handler.HttpRequest,
		order.NewOrderGetAll(orderSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/orders", rtr.handle(
		handler.HttpRequest,
		order.NewOrderCreate(orderSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/orders/{order_id}", rtr.handle(
		handler.HttpRequest,
		order.NewOrderGetByID(orderSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/orders/{order_id}/payment", rtr.handle(
		handler.HttpRequest,
		order.NewOrderPayment(orderSvc),
	)).Methods(http.MethodPost)

}
