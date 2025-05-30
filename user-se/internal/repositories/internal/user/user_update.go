package user

import (
	"auth-se/internal/presentations"
	"auth-se/pkg/postgres"
	"context"
	"database/sql"
	"time"

	"auth-se/internal/consts"
	"github.com/pkg/errors"
)

func (c user) UpdateUser(ctx context.Context, userID string, input presentations.UserUpdate) error {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return errors.Wrap(err, "failed begin tx")
	}

	query := `
	UPDATE users 
	SET 
	    name = $2, 
	    email = $3, 
	    phone = $4, 
	    role_id = $5, 
	    updated_at = $6 
	WHERE id = $1 
	AND deleted_at is null;`

	values := []interface{}{
		userID,
		input.Name,
		input.Email,
		input.Phone,
		input.RoleID,
		time.Now().Local(),
	}

	if _, err := c.db.ExecTx(ctx, tx, query, values...); err != nil {
		errRollback := c.db.RollbackTx(ctx, tx)
		if errRollback != nil {
			return errors.Wrap(err, "rollback failed")
		}

		errSql := c.db.ParseSQLError(err)

		if errSql != nil {
			switch errSql {
			case postgres.ErrUniqueViolation:
				return consts.ErrPhoneAlreadyExist

			default:
				return errors.Wrap(err, "failed execute query")
			}
		}

	}

	if err := c.db.CommitTx(ctx, tx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}
