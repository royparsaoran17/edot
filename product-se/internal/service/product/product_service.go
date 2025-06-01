package product

import (
	"context"
	"github.com/google/uuid"
	"product-se/internal/common"

	"github.com/pkg/errors"
	"product-se/internal/entity"
	"product-se/internal/presentations"
	"product-se/internal/repositories"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) Product {
	return &service{repo: repo}
}

func (s *service) GetAllProduct(ctx context.Context, meta *common.Metadata) ([]entity.Product, error) {
	products, err := s.repo.Product.GetAllProduct(ctx, meta)
	if err != nil {
		return nil, errors.Wrap(err, "getting all products on ")
	}

	return products, nil
}

func (s *service) GetProductByID(ctx context.Context, productID string) (*entity.Product, error) {
	products, err := s.repo.Product.FindProductByID(ctx, productID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting product id %s", productID)
	}

	return products, nil
}

func (s *service) UpdateProductByID(ctx context.Context, productID string, input presentations.ProductUpdate) error {
	if err := input.Validate(); err != nil {
		return errors.Wrap(err, "validation(s) error")
	}

	_, err := s.repo.Product.FindProductByID(ctx, productID)
	if err != nil {
		return errors.Wrapf(err, "getting product id %s", productID)
	}

	if err := s.repo.Product.UpdateProduct(ctx, productID, input); err != nil {
		return errors.Wrap(err, "updating product")

	}

	return nil
}

func (s *service) CreateProduct(ctx context.Context, input presentations.ProductCreate) (*entity.Product, error) {
	if err := input.Validate(); err != nil {
		return nil, errors.Wrap(err, "validation(s) error")
	}

	productID := uuid.NewString()
	input.ID = productID
	err := s.repo.Product.CreateProduct(ctx, input)
	if err != nil {
		return nil, errors.Wrap(err, "creating product")

	}

	product, err := s.repo.Product.FindProductByID(ctx, productID)
	if err != nil {
		return nil, errors.Wrapf(err, "getting product id %s", productID)
	}

	return product, nil
}
