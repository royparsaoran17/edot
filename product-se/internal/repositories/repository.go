package repositories

import (
	"product-se/internal/repositories/internal/product"
	"product-se/internal/repositories/internal/warehouse"
	"product-se/pkg/postgres"
)

type Repository struct {
	Product   Product
	Warehouse Warehouse
}

func NewRepository(db postgres.Adapter) *Repository {
	return &Repository{
		Product:   product.NewProduct(db),
		Warehouse: warehouse.NewWarehouse(db),
	}
}
