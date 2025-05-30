package warehouse

import (
	"context"
	"database/sql"
	"product-se/internal/presentations"
	"product-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
	"product-se/internal/consts"
)

func (r warehouse) UpdateWarehouse(ctx context.Context, warehouseID string, input presentations.WarehouseUpdate) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed begin tx")
	}

	query := `
	UPDATE warehouses 
	SET 
	    name = $2, 
	    updated_at = $3 
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		warehouseID,
		input.Name,
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
				return consts.ErrWarehouseAlreadyExist

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
