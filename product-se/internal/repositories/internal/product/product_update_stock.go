package product

import (
	"context"
	"database/sql"
	"product-se/internal/consts"
	"product-se/internal/presentations"
	"product-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
)

func (r product) UpdateProductStock(ctx context.Context, stockID string, input presentations.ProductUpdateStock) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed begin tx")
	}

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

	if _, err := r.db.ExecTx(ctx, tx, query, values...); err != nil {
		errRollback := r.db.RollbackTx(ctx, tx)
		if errRollback != nil {
			return errors.Wrap(err, "rollback failed")
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

	if err := r.db.CommitTx(ctx, tx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
