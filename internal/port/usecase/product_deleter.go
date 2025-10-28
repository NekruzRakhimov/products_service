package usecase

import "context"

type ProductDeleter interface {
	DeleteProductByID(ctx context.Context, id int) (err error)
}
