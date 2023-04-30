package service

import (
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/repository/product_repository"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductService_GetProductById_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)
	currentTime := time.Now()
	product := entity.Product{
		Id:          1,
		Title:       "Test Product",
		Description: "Description of product",
		Price:       1000,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
	product_repository.GetProductById = func(productId int) (*entity.Product, errs.MessageErr) {
		return &product, nil
	}
	response, err := productService.GetProductById(1)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Test Product", response.Title)
	assert.Equal(t, 1, response.Id)
}
func TestProductService_GetProductId_NotFoundError(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)
	product_repository.GetProductById = func(productId int) (*entity.Product, errs.MessageErr) {
		return nil, errs.NewNotFoundError("product data not found")
	}
	response, err := productService.GetProductById(1)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, "product data not found", err.Message())
	assert.Equal(t, "NOT_FOUND", err.Error())
}
func TestProductService_GetProduct_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)
	currentTime := time.Now()
	products := []*entity.Product{
		{
			Id:          1,
			Title:       "Test Product",
			Description: "Product description of Test Product",
			Price:       2000,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
	}
	product_repository.GetProduct = func(userId int) ([]*entity.Product, errs.MessageErr) {
		return products, nil
	}
	response, err := productService.GetProduct(1)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "Test Product", response.Data[0].Title)
}
func TestTestProductService_GetProduct_NotFound(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)
	product_repository.GetProduct = func(userId int) ([]*entity.Product, errs.MessageErr) {
		return []*entity.Product{}, nil
	}
	response, err := productService.GetProduct(1)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 0, len(response.Data))
}
