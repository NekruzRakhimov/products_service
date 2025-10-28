package dbstore

import (
	"github.com/jmoiron/sqlx"
)

type DBStore struct {
	ProductStorage *ProductStorage
}

func New(db *sqlx.DB) *DBStore {
	return &DBStore{
		ProductStorage: NewProductStorage(db),
	}
}
