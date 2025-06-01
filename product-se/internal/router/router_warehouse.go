package router

import (
	"net/http"
	"product-se/internal/handler"
	"product-se/internal/ucase/warehouse"

	warehousesvc "product-se/internal/service/warehouse"
)

func (rtr *router) mountWarehouses(warehouseSvc warehousesvc.Warehouse) {
	rtr.router.HandleFunc("/internal/v1/warehouses", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseGetAll(warehouseSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/warehouses", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseCreate(warehouseSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseGetByID(warehouseSvc),
	)).Methods(http.MethodGet)

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseUpdate(warehouseSvc),
	)).Methods(http.MethodPut)

}
