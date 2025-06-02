package router

import (
	"net/http"
	"warehouse-se/internal/handler"
	"warehouse-se/internal/ucase/warehouse"

	warehousesvc "warehouse-se/internal/service/warehouse"
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

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}/add-stock", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseAddStock(warehouseSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}/deduct-stock", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseDeductStock(warehouseSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}/move-stock", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseMoveStock(warehouseSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}/activate", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseActivate(warehouseSvc),
	)).Methods(http.MethodPost)

	rtr.router.HandleFunc("/internal/v1/warehouses/{warehouse_id}/inactive", rtr.handle(
		handler.HttpRequest,
		warehouse.NewWarehouseInactive(warehouseSvc),
	)).Methods(http.MethodPost)

}
