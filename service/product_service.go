package service

import (
	"chapter3-sesi2/dto"
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/pkg/helpers"
	"chapter3-sesi2/repository/product_repository"
	"fmt"
	"net/http"
)

type ProductService interface {
	CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	GetProduct(userId int) (*dto.GetProductResponse, errs.MessageErr)
	GetProductById(productId int) (*dto.ProductResponse, errs.MessageErr)
	UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	DeleteProductById(productId int) (*dto.NewProductResponse, errs.MessageErr)
}
type productService struct {
	productRepo product_repository.ProductRepository
}

// GetProductById implements ProductService
func (p *productService) GetProductById(productId int) (*dto.ProductResponse, errs.MessageErr) {
	result, err := p.productRepo.GetProductById(productId)
	if err != nil {
		return nil, err
	}
	response := result.EntityToProductResponseDto()
	return &response, nil
}

// GetProduct implements ProductService
func (p *productService) GetProduct(userId int) (*dto.GetProductResponse, errs.MessageErr) {
	products, err := p.productRepo.GetProduct(userId)
	if err != nil {
		return nil, err
	}
	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponse := product.EntityToProductResponseDto()
		productResponses = append(productResponses, productResponse)
	}
	response := dto.GetProductResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "Your products successfully loaded",
		Data:       productResponses,
	}
	return &response, nil
}

// CreateProduct implements ProductService
func (p *productService) CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	productRequest := &entity.Product{
		Title:       payload.Title,
		Description: payload.Description,
		Price:       payload.Price,
		UserId:      userId,
	}
	_, err := p.productRepo.CreateProduct(productRequest)
	if err != nil {
		return nil, err
	}
	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new product data sucessfully created",
	}
	return &response, err
}

// UpdateProductById implements ProductService
func (p *productService) UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(productRequest)
	if err != nil {
		return nil, err
	}
	payload := entity.Product{
		Id:          productId,
		Title:       productRequest.Title,
		Description: productRequest.Description,
		Price:       productRequest.Price,
	}
	err = p.productRepo.UpdateProductById(payload)
	if err != nil {
		return nil, err
	}
	response := dto.NewProductResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "product data successfully updated",
	}
	return &response, nil
}
func (p *productService) DeleteProductById(productId int) (*dto.NewProductResponse, errs.MessageErr) {
	err := p.productRepo.DeleteProductById(productId)
	if err != nil {
		return nil, err
	}
	message := fmt.Sprintf("product data with id %d has been deleted", productId)
	response := dto.NewProductResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    message,
	}
	return &response, nil
}

func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}
