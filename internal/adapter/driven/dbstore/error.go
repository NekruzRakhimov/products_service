package dbstore

import (
	"database/sql"
	"errors"
	"github.com/NekruzRakhimov/product_service/internal/errs"
)

func (u *ProductStorage) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}
