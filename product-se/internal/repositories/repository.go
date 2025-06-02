package repositories

import (
	"context"
	"database/sql"
	"product-se/internal/repositories/internal/product"
	"product-se/pkg/databasex"
)

type Repository struct {
	Product Product
	db      databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{
		Product: product.NewProduct(db),
	}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
