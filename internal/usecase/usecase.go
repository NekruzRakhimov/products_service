package usecase

import (
	"github.com/NekruzRakhimov/product_service/internal/adapter/driven/dbstore"
	"github.com/NekruzRakhimov/product_service/internal/config"
	"github.com/NekruzRakhimov/product_service/internal/port/driven"
	"github.com/NekruzRakhimov/product_service/internal/port/usecase"
	productcreater "github.com/NekruzRakhimov/product_service/internal/usecase/product_creater"
	productdeleter "github.com/NekruzRakhimov/product_service/internal/usecase/product_deleter"
	productgetter "github.com/NekruzRakhimov/product_service/internal/usecase/product_getter"
	productupdater "github.com/NekruzRakhimov/product_service/internal/usecase/product_updater"
)

type UseCases struct {
	ProductCreater usecase.ProductCreater
	ProductUpdater usecase.ProductUpdater
	ProductDeleter usecase.ProductDeleter
	ProductGetter  usecase.ProductGetter
}

func New(cfg *config.Config, store *dbstore.DBStore, cache driven.Cache) *UseCases {
	return &UseCases{
		ProductGetter:  productgetter.New(cfg, store.ProductStorage, cache),
		ProductCreater: productcreater.New(cfg, store.ProductStorage),
		ProductUpdater: productupdater.New(cfg, store.ProductStorage),
		ProductDeleter: productdeleter.New(cfg, store.ProductStorage),
	}
}
