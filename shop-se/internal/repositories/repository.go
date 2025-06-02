package repositories

import (
	"context"
	"database/sql"
	"shop-se/internal/repositories/internal/shop"
	"shop-se/pkg/databasex"
)

type Repository struct {
	Shop Shop
	db   databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Shop: shop.NewShop(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
