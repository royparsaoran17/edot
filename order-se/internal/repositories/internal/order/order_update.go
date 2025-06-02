package order

import (
	"context"
	"database/sql"
	"order-se/internal/presentations"
	"order-se/internal/repositories/repooption"
	"order-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
	"order-se/internal/consts"
)

func (r order) UpdateOrder(ctx context.Context, orderID string, input presentations.OrderUpdate, opts ...repooption.TxOption) error {

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
	UPDATE orders 
	SET 
	    status = $2, 
	    updated_at = $3
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		orderID,
		input.Status,
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
				return consts.ErrOrderAlreadyExist

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
