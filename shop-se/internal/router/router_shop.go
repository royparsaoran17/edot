package router

import (
	"net/http"
	"shop-se/internal/handler"
	"shop-se/internal/ucase/shop"

	shopsvc "shop-se/internal/service/shop"
)

func (rtr *router) mountShops(shopSvc shopsvc.Shop) {
	rtr.router.HandleFunc("/internal/v1/shops", rtr.handle(
		handler.HttpRequest,
		shop.NewShopGetAll(shopSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/shops", rtr.handle(
		handler.HttpRequest,
		shop.NewShopCreate(shopSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/shops/{shop_id}", rtr.handle(
		handler.HttpRequest,
		shop.NewShopGetByID(shopSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/shops/{shop_id}", rtr.handle(
		handler.HttpRequest,
		shop.NewShopUpdate(shopSvc),
	)).Methods(http.MethodPut)

}
