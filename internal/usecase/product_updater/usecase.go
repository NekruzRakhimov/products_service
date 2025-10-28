package product_updater

import (
	"context"
	"errors"
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

func (u *UseCase) UpdateProductByID(ctx context.Context, product domain.Product) (err error) {
	_, err = u.productStorage.GetProductByID(ctx, product.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrProductNotfound
		}
		return err
	}

	err = u.productStorage.UpdateProductByID(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
