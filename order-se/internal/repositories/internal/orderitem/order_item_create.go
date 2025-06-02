package orderitem

import (
	"context"
	"database/sql"
	"order-se/internal/repositories/repooption"
	"time"

	"order-se/internal/consts"
	"order-se/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"order-se/internal/presentations"
)

func (r orderItem) CreateOrderItem(ctx context.Context, input presentations.OrderItemCreate, opts ...repooption.TxOption) error {

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
	INSERT INTO order_items (
        id, 
		order_id, 
		product_id, 
		warehouse_id, 
		quantity, 
		price, 
		created_at, 
		updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $7)`

	values := []interface{}{
		uuid.New().String(),
		input.OrderID,
		input.Price,
		input.WarehouseID,
		input.Quantity,
		input.Price,
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
				return consts.ErrDataAlreadyExist

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
