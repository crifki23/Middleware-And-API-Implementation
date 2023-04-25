package service

import (
	"chapter3-sesi2/dto"
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/pkg/helpers"
	"chapter3-sesi2/repository/product_repository"
	"net/http"
)

type ProductService interface {
	CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
}
type productService struct {
	productRepo product_repository.ProductRepository
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

func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}
