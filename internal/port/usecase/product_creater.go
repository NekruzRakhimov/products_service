package usecase

import (
	"context"
	"github.com/NekruzRakhimov/product_service/internal/domain"
)

type ProductCreater interface {
	CreateProduct(ctx context.Context, product domain.Product) (err error)
}
