package repositories

import (
	"context"
	"database/sql"
	"shop-se/internal/repositories/internal/shop"
	"shop-se/internal/repositories/internal/user"
	"shop-se/internal/repositories/internal/warehouse"
	"shop-se/pkg/databasex"
)

type Repository struct {
	Shop      Shop
	User      User
	Warehouse Warehouse
	db        databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Shop:      shop.NewShop(db),
		User:      user.NewUser(db),
		Warehouse: warehouse.NewWarehouse(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
