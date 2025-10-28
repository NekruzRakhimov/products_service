package usecase

import (
	"context"
	"github.com/NekruzRakhimov/product_service/internal/domain"
)

type ProductGetter interface {
	GetAllProducts(ctx context.Context) (products []domain.Product, err error)
	GetProductByID(ctx context.Context, id int) (product domain.Product, err error)
}
