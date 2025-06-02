package shop

import (
	"context"
	"database/sql"
	"shop-se/internal/repositories/repooption"
	"time"

	"shop-se/internal/consts"
	"shop-se/pkg/postgres"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"shop-se/internal/presentations"
)

func (r shop) CreateShop(ctx context.Context, input presentations.ShopCreate, opts ...repooption.TxOption) error {

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

	query := `INSERT INTO shops (id, name, owner_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $4)`

	values := []interface{}{
		uuid.New().String(),
		input.Name,
		input.OwnerID,
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
				return consts.ErrShopAlreadyExist

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
