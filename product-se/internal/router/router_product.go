package router

import (
	"net/http"
	"product-se/internal/handler"
	"product-se/internal/ucase/product"

	productsvc "product-se/internal/service/product"
)

func (rtr *router) mountProducts(productSvc productsvc.Product) {
	rtr.router.HandleFunc("/internal/v1/products", rtr.handle(
		handler.HttpRequest,
		product.NewProductGetAll(productSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/products", rtr.handle(
		handler.HttpRequest,
		product.NewProductCreate(productSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/products/{product_id}", rtr.handle(
		handler.HttpRequest,
		product.NewProductGetByID(productSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/products/{product_id}", rtr.handle(
		handler.HttpRequest,
		product.NewProductUpdate(productSvc),
	)).Methods(http.MethodPut)

}
