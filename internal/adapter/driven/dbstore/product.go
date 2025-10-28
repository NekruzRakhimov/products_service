package dbstore

import (
	"context"
	"github.com/NekruzRakhimov/product_service/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"os"
	"time"
)

type ProductStorage struct {
	db *sqlx.DB
}

func NewProductStorage(db *sqlx.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

type Product struct {
	ID           int       `db:"id"`
	ProductName  string    `db:"product_name"`
	Manufacturer string    `db:"manufacturer"`
	ProductCount int       `db:"product_count"`
	Price        float64   `db:"price"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (p *Product) FromDomain(dProduct domain.Product) {
	p.ID = dProduct.ID
	p.ProductName = dProduct.ProductName
	p.Manufacturer = dProduct.Manufacturer
	p.ProductCount = dProduct.ProductCount
	p.Price = dProduct.Price
	p.CreatedAt = dProduct.CreatedAt
	p.UpdatedAt = dProduct.UpdatedAt
}

func (p *Product) ToDomain() domain.Product {
	return domain.Product{
		ID:           p.ID,
		ProductName:  p.ProductName,
		Manufacturer: p.Manufacturer,
		ProductCount: p.ProductCount,
		Price:        p.Price,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func (r *ProductStorage) GetAllProducts(ctx context.Context) (products []domain.Product, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAllProducts").Logger()

	var dbProducts []Product
	if err = r.db.SelectContext(ctx, &dbProducts, `
		SELECT id, product_name, manufacturer, product_count, price, created_at, updated_at
		FROM products
		ORDER BY id`); err != nil {
		logger.Err(err).Msg("error selecting products")
		return nil, r.translateError(err)
	}

	for _, dbProduct := range dbProducts {
		products = append(products, dbProduct.ToDomain())
	}

	return products, nil
}

func (r *ProductStorage) GetProductByID(ctx context.Context, id int) (domain.Product, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetProductByID").Logger()

	var dbProduct Product
	if err := r.db.GetContext(ctx, &dbProduct, `
		SELECT id, product_name, manufacturer, product_count, price, created_at, updated_at
		FROM products
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting product")
		return domain.Product{}, r.translateError(err)
	}

	return dbProduct.ToDomain(), nil
}

func (r *ProductStorage) CreateProduct(ctx context.Context, product domain.Product) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateProduct").Logger()

	var dbProduct Product
	dbProduct.FromDomain(product)
	_, err = r.db.ExecContext(ctx, `INSERT INTO products (product_name, manufacturer, price, product_count)
					VALUES ($1, $2, $3, $4)`,
		dbProduct.ProductName,
		dbProduct.Manufacturer,
		dbProduct.Price,
		dbProduct.ProductCount)
	if err != nil {
		logger.Err(err).Msg("error inserting product")
		return r.translateError(err)
	}

	return nil
}

func (r *ProductStorage) UpdateProductByID(ctx context.Context, product domain.Product) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.UpdateProductByID").Logger()
	var dbProduct Product
	dbProduct.FromDomain(product)
	_, err = r.db.ExecContext(ctx, `
		UPDATE products SET product_name = $1, 
		                    manufacturer = $2, 
		                    price = $3,
		                    product_count = $4
		                WHERE id = $5`,
		dbProduct.ProductName,
		dbProduct.Manufacturer,
		dbProduct.Price,
		dbProduct.ProductCount,
		dbProduct.ID)
	if err != nil {
		logger.Err(err).Msg("error updating product")
		return r.translateError(err)
	}

	return nil
}

func (r *ProductStorage) DeleteProductByID(ctx context.Context, id int) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.DeleteProductByID").Logger()
	_, err = r.db.ExecContext(ctx, `DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting product")
		return r.translateError(err)
	}

	return nil
}
