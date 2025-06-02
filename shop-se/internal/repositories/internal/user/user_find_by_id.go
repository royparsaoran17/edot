package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
	"shop-se/internal/consts"
	"shop-se/internal/entity"
)

func (c user) FindUserByID(ctx context.Context, userID string) (*entity.User, error) {
	query := `SELECT 
        jsonb_build_object(
            'id', c.id,
            'name', c.name,
            'email', c.email,
            'phone', c.phone,
            'password', c.password,
            'role_id', c.role_id,
            'created_at', c.created_at::timestamptz,
            'updated_at', c.updated_at::timestamptz,
            'deleted_at', c.deleted_at::timestamptz
        )
    FROM
        users c
    WHERE c.id = $1
        AND c.deleted_at is null;`

	var b []byte
	err := c.db.QueryRow(ctx, query, userID).Scan(&b)
	if err != nil {
		sqlErr := c.db.ParseSQLError(err)
		switch sqlErr {
		case sql.ErrNoRows:
			return nil, consts.ErrUserNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch user from db")
		}
	}

	var role entity.User
	if err := json.Unmarshal(b, &role); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal byte to user")
	}

	return &role, nil
}
