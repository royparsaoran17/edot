package warehouse

import (
	"context"
	"database/sql"
	"warehouse-se/internal/presentations"
	"warehouse-se/internal/repositories/repooption"
	"warehouse-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
	"warehouse-se/internal/consts"
)

func (r warehouse) UpdateWarehouse(ctx context.Context, warehouseID string, input presentations.WarehouseUpdate, opts ...repooption.TxOption) error {

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
	UPDATE warehouses 
	SET 
	    name = $2, 
	    is_active = $3,
	    updated_at = $4 
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		warehouseID,
		input.Name,
		input.IsActive,
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
				return consts.ErrWarehouseAlreadyExist

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
