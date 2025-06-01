package product

import (
	"context"
	"database/sql"
	"product-se/internal/consts"
	"product-se/internal/presentations"
	"product-se/internal/repositories/repooption"
	"product-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
)

func (r product) UpdateProductStock(ctx context.Context, stockID string, input presentations.ProductUpdateStock, opts ...repooption.TxOption) error {

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
	query := `
	UPDATE products 
	SET 
	    quantity = $2, 
	    updated_at = $3 
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		stockID,
		input.Quantity,
		time.Now().Local(),
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
