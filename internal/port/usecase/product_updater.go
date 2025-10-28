package usecase

import (
	"context"
	"github.com/NekruzRakhimov/product_service/internal/domain"
)

type ProductUpdater interface {
	UpdateProductByID(ctx context.Context, product domain.Product) (err error)
}
