package user

import (
	"auth-se/internal/consts"
	"auth-se/internal/entity"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/pkg/errors"
)

func (c user) FindUserByPhoneEmailPassword(ctx context.Context, value, password string) (*entity.UserDetail, error) {
	query := `SELECT 
        jsonb_build_object(
            'id', c.id,
            'name', c.name,
            'phone', c.phone,
            'email', c.email,
            'password', c.password,
            'role_id', c.role_id,
            'role',(
                SELECT
					json_build_object(
						'id', r.id,
						'name', r.name,
						'created_at', r.created_at::timestamptz,
						'updated_at', r.updated_at::timestamptz,
						'deleted_at', r.deleted_at::timestamptz
					)
                FROM roles r
                    WHERE c.role_id = r.id
                    AND r.deleted_at is null
                ),
            'created_at', c.created_at::timestamptz,
            'updated_at', c.updated_at::timestamptz,
            'deleted_at', c.deleted_at::timestamptz
        )
    FROM
        users c
    WHERE (c.phone = $1 or c.email = $1)
    AND c.password = $2
        AND c.deleted_at is null;`

	var b []byte
	err := c.db.QueryRow(ctx, query, value, password).Scan(&b)
	if err != nil {
		sqlErr := c.db.ParseSQLError(err)
		switch sqlErr {
		case sql.ErrNoRows:
			return nil, consts.ErrUserNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch user from db")
		}
	}

	var role entity.UserDetail
	if err := json.Unmarshal(b, &role); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal byte to user")
	}

	return &role, nil
}
