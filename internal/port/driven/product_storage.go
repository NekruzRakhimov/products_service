package driven

import (
	"context"

	"github.com/NekruzRakhimov/product_service/internal/domain"
)

type ProductStorage interface {
	GetAllProducts(ctx context.Context) (products []domain.Product, err error)
	GetProductByID(ctx context.Context, id int) (domain.Product, error)
	CreateProduct(ctx context.Context, product domain.Product) (err error)
	UpdateProductByID(ctx context.Context, product domain.Product) (err error)
	DeleteProductByID(ctx context.Context, id int) (err error)
}
