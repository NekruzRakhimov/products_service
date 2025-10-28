package product_getter

import (
	"context"
	"errors"
	"fmt"
	"github.com/NekruzRakhimov/product_service/internal/config"
	"github.com/NekruzRakhimov/product_service/internal/domain"
	"github.com/NekruzRakhimov/product_service/internal/errs"
	"github.com/NekruzRakhimov/product_service/internal/port/driven"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	defaultTTL = time.Minute * 5
)

type UseCase struct {
	cfg            *config.Config
	productStorage driven.ProductStorage
	cache          driven.Cache
}

func New(cfg *config.Config, productStorage driven.ProductStorage, cache driven.Cache) *UseCase {
	return &UseCase{
		cfg:            cfg,
		productStorage: productStorage,
		cache:          cache,
	}
}

func (u *UseCase) GetAllProducts(ctx context.Context) (products []domain.Product, err error) {
	products, err = u.productStorage.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *UseCase) GetProductByID(ctx context.Context, id int) (product domain.Product, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "service.GetProductByID").Logger()

	// 1. Обращаемся в кэш (redis) и проверяем есть ли продукт с таким id
	err = u.cache.Get(ctx, fmt.Sprintf("product_%d", id), &product)
	if err == nil {
		// 2. Если есть, то возвращаем клиенту
		return product, nil
	}

	if !errors.Is(err, redis.Nil) {
		return domain.Product{}, err
	}

	// 3. Если нет, то обращаемся в бд (postgres)
	product, err = u.productStorage.GetProductByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return domain.Product{}, errs.ErrProductNotfound
		}
		return domain.Product{}, err
	}

	// 3.1 Сохраняем в кэш (redis)
	if err = u.cache.Set(ctx, fmt.Sprintf("product_%d", product.ID), product, defaultTTL); err != nil {
		logger.Error().Err(err).Int("product_id", product.ID).Msg("error setting product in cache")
	}

	// 3.2 Возвращаем клиенту
	return product, nil

}
