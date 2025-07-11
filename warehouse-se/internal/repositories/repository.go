package repositories

import (
	"context"
	"database/sql"
	"warehouse-se/internal/repositories/internal/product"
	"warehouse-se/internal/repositories/internal/warehouse"
	"warehouse-se/pkg/databasex"
)

type Repository struct {
	Product   Product
	Warehouse Warehouse
	db        databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Product:   product.NewProduct(db),
		Warehouse: warehouse.NewWarehouse(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
