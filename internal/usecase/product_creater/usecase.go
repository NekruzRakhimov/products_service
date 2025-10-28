package product_creater

import (
	"context"
	"github.com/NekruzRakhimov/product_service/internal/config"
	"github.com/NekruzRakhimov/product_service/internal/domain"
	"github.com/NekruzRakhimov/product_service/internal/errs"
	"github.com/NekruzRakhimov/product_service/internal/port/driven"
)

type UseCase struct {
	cfg            *config.Config
	productStorage driven.ProductStorage
}

func New(cfg *config.Config, productStorage driven.ProductStorage) *UseCase {
	return &UseCase{
		cfg:            cfg,
		productStorage: productStorage,
	}
}

func (u *UseCase) CreateProduct(ctx context.Context, product domain.Product) (err error) {
	if len(product.ProductName) < 4 {
		return errs.ErrInvalidProductName
	}

	err = u.productStorage.CreateProduct(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
