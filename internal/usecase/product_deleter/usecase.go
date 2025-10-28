package product_deleter

import (
	"context"
	"errors"
	"github.com/NekruzRakhimov/product_service/internal/config"
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

func (u *UseCase) DeleteProductByID(ctx context.Context, id int) (err error) {
	_, err = u.productStorage.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrProductNotfound
		}
		return err
	}

	err = u.productStorage.DeleteProductByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
