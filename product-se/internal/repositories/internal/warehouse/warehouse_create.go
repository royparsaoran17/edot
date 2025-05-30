package warehouse

import (
	"context"
	"database/sql"
	"time"

	"product-se/internal/consts"
	"product-se/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"product-se/internal/presentations"
)

func (r warehouse) CreateWarehouse(ctx context.Context, input presentations.WarehouseCreate) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed begin tx")
	}

	tNow := time.Now().Local()

	query := `INSERT INTO warehouses (id, name, created_at, updated_at) VALUES ($1, $2, $3, $3)`

	values := []interface{}{
		uuid.New().String(),
		input.Name,
		tNow,
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
