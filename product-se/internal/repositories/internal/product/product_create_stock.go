package product

import (
	"context"
	"database/sql"
	"product-se/internal/repositories/repooption"
	"time"

	"product-se/internal/consts"
	"product-se/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"product-se/internal/presentations"
)

func (r product) CreateProductStock(ctx context.Context, input presentations.ProductCreateStock, opts ...repooption.TxOption) error {

	txOpt := repooption.TxOptions{
		Tx:              nil,
		NotCommitInRepo: false,
	}

	for _, opt := range opts {
		opt(&txOpt)
	}

	if txOpt.Tx == nil {
		tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
		if err != nil {
			return errors.Wrap(err, "beginTx")
		}

		txOpt.Tx = tx
	}

	tx := txOpt.Tx
	tNow := time.Now().Local()

	query := `
		INSERT INTO products (
			id, 
			product_id, 
			warehouse_id,
			quantity,
			created_at, 
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $5)`

	values := []interface{}{
		uuid.New().String(),
		input.ProductID,
		input.WarehouseID,
		input.Quantity,
		tNow,
	}

	if _, err := tx.ExecContext(ctx, query, values...); err != nil {
		if !txOpt.NotCommitInRepo {
			if err := tx.Rollback(); err != nil {
				err = errors.Wrap(err, "rollback")
			}
		}
		errSql := r.db.ParseSQLError(err)

		if errSql != nil {
			switch errSql {
			case postgres.ErrUniqueViolation:
				return consts.ErrProductAlreadyExist

			default:
				return errors.Wrap(err, "failed execute query")
			}
		}

	}

	if !txOpt.NotCommitInRepo {
		err := tx.Commit()
		if err != nil {
			return errors.Wrap(err, "commit add chopper")
		}
	}
	return nil
}
