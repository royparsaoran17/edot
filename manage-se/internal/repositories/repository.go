package repositories

import (
	"context"
	"database/sql"
	"manage-se/pkg/databasex"
)

type Repository struct {
	db databasex.Adapter
}

func NewRepository(db databasex.Adapter) *Repository {
	return &Repository{}
}

func (r Repository) BeginTx(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, options)
}
