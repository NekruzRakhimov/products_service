package http

import (
	"errors"
	"github.com/NekruzRakhimov/product_service/internal/domain"
	"github.com/NekruzRakhimov/product_service/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Product struct {
	ID           int       `json:"id"`
	ProductName  string    `json:"product_name"`
	Manufacturer string    `json:"manufacturer"`
	ProductCount int       `json:"product_count"`
	Price        float64   `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (p *Product) fromDomain(dProduct domain.Product) {
	p.ID = dProduct.ID
	p.ProductName = dProduct.ProductName
	p.Manufacturer = dProduct.Manufacturer
	p.ProductCount = dProduct.ProductCount
	p.Price = dProduct.Price
	p.CreatedAt = dProduct.CreatedAt
	p.UpdatedAt = dProduct.UpdatedAt
}

// GetAllProducts
// @Summary Получение продуктов
// @Description Получение списка всех продуктов
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {array} Product
// @Failure 500 {object} CommonError
// @Router /api/products [get]
func (s *Server) GetAllProducts(c *gin.Context) {
	logger := zerolog.New(os.Stdout).With().Str("func_name", "controller.GetAllProducts").Logger()
	userID := c.GetInt(userIDCtx)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, CommonError{Error: "invalid userID in context"})
		return
	}

	logger.Debug().Int("user_id", userID).Msg("GetUser")

	dProducts, err := s.uc.ProductGetter.GetAllProducts(c)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var (
		products []Product
		product  Product
	)
	for _, dProduct := range dProducts {
		product.fromDomain(dProduct)
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID
// @Summary Получить продукт по ID
// @Description Получение продукта по ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "id продукта"
// @Success 200 {object} Product
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/products/{id} [get]
func (s *Server) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidProductID)
		return
	}

	dProduct, err := s.uc.ProductGetter.GetProductByID(c, id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	var product Product
	product.fromDomain(dProduct)

	c.JSON(http.StatusOK, product)
}

type CreateProductRequest struct {
	ProductName  string  `json:"product_name"`
	Manufacturer string  `json:"manufacturer"`
	ProductCount int     `json:"product_count"`
	Price        float64 `json:"price"`
}

// CreateProduct
// @Summary Добавление нового продукта
// @Description Добавление нового продукта
// @Tags Products
// @Consume json
// @Produce json
// @Security BearerAuth
// @Param request_body body CreateProductRequest true "информация о новом продукте"
// @Success 201 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/products [post]
func (s *Server) CreateProduct(c *gin.Context) {
	var product CreateProductRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		s.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if product.ProductName == "" || product.Manufacturer == "" || product.ProductCount < 0 || product.Price < 0 {
		s.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err := s.uc.ProductCreater.CreateProduct(c, domain.Product{
		ProductName:  product.ProductName,
		Manufacturer: product.Manufacturer,
		ProductCount: product.ProductCount,
		Price:        product.Price,
	}); err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CommonResponse{Message: "Product created successfully"})
}

// UpdateProductByID
// @Summary Обновить продукт по ID
// @Description Обновление продукта по ID
// @Tags Products
// @Consume json
// @Produce json
// @Security BearerAuth
// @Param id path int true "id продукта"
// @Param request_body body CreateProductRequest true "информация о продукте"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 422 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/products/{id} [put]
func (s *Server) UpdateProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidProductID)
		return
	}

	var product CreateProductRequest
	if err = c.ShouldBindJSON(&product); err != nil {
		s.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if product.ProductName == "" || product.Manufacturer == "" || product.ProductCount < 0 || product.Price < 0 {
		s.handleError(c, errs.ErrInvalidFieldValue)
		return
	}

	if err = s.uc.ProductUpdater.UpdateProductByID(c, domain.Product{
		ID:           id,
		ProductName:  product.ProductName,
		Manufacturer: product.Manufacturer,
		ProductCount: product.ProductCount,
		Price:        product.Price,
	}); err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Product updated successfully"})
}

// DeleteProductByID
// @Summary Удалить продукт по ID
// @Description Удаление продукта по ID
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param id path int true "id продукта"
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonError
// @Failure 404 {object} CommonError
// @Failure 500 {object} CommonError
// @Router /api/products/{id} [delete]
func (s *Server) DeleteProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		s.handleError(c, errs.ErrInvalidProductID)
		return
	}

	if err = s.uc.ProductDeleter.DeleteProductByID(c, id); err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, CommonResponse{Message: "Product deleted successfully"})
}
