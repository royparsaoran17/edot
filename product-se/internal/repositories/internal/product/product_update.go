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

func (r product) UpdateProduct(ctx context.Context, productID string, input presentations.ProductUpdate) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed begin tx")
	}

	query := `
	UPDATE products 
	SET 
	    name = $2, 
	    description = $3, 
	    price = $4, 
	    unit = $5, 
	    sku = $6, 
	    category = $7, 
	    is_active = $8, 
	    updated_at = $9 
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		productID,
		input.Name,
		input.Description,
		input.Price,
		input.Unit,
		input.Sku,
		input.Category,
		input.IsActive,
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
