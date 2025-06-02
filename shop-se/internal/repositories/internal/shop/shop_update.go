package shop

import (
	"context"
	"database/sql"
	"shop-se/internal/presentations"
	"shop-se/internal/repositories/repooption"
	"shop-se/pkg/postgres"
	"time"

	"github.com/pkg/errors"
	"shop-se/internal/consts"
)

func (r shop) UpdateShop(ctx context.Context, shopID string, input presentations.ShopUpdate, opts ...repooption.TxOption) error {

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
	UPDATE shops 
	SET 
	    name = $2, 
	    owner_id = $3,
	    updated_at = $4 
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		shopID,
		input.Name,
		input.OwnerID,
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
