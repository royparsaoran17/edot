package product

import (
	"context"
	"database/sql"
	"order-se/internal/consts"
	"order-se/internal/entity"

	"github.com/pkg/errors"
)

func (r product) FindProductByID(ctx context.Context, productID string) (*entity.Product, error) {
	query := `
		SELECT 
			id,
			name,
			description,
			price,
			unit,
			sku,
			category,
			is_active,
			created_at::timestamptz,
			updated_at::timestamptz, 
			deleted_at::timestamptz
		FROM products 
		WHERE id = $1
		  AND deleted_at is null
		  AND is_active = 1
`

	var product entity.Product

	err := r.db.QueryRow(ctx, &product, query, productID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, consts.ErrProductNotFound
		default:
			return nil, errors.Wrap(err, "failed to fetch row from db")
		}
	}

	return &product, nil
}
