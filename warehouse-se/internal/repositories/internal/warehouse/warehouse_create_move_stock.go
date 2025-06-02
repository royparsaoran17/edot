package warehouse

import (
	"context"
	"database/sql"
	"warehouse-se/internal/repositories/repooption"
	"time"

	"warehouse-se/internal/consts"
	"warehouse-se/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"warehouse-se/internal/presentations"
)

func (r warehouse) CreateMoveStockWarehouse(ctx context.Context, input presentations.WarehouseCreateMoveStock, opts ...repooption.TxOption) error {

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

	query := `INSERT INTO stock_movements (id, product_id, from_warehouse_id, to_warehouse_id, quantity, notes, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $7)`

	values := []interface{}{
		uuid.New().String(),
		input.ProductID,
		input.FromWarehouseID,
		input.ToWarehouseID,
		input.Quantity,
		input.Notes,
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
