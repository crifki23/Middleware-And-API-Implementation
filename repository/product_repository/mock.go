package product_repository

import (
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
)

var (
	CreateProduct     func(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById    func(productId int) (*entity.Product, errs.MessageErr)
	GetProduct        func(userId int) ([]*entity.Product, errs.MessageErr)
	UpdateProductById func(payload entity.Product) errs.MessageErr
	DeleteProductById func(productId int) errs.MessageErr
)

type productRepoMock struct{}

// CreateProduct implements ProductRepository
func (*productRepoMock) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	return CreateProduct(productPayload)
}

// DeleteProductById implements ProductRepository
func (*productRepoMock) DeleteProductById(productId int) errs.MessageErr {
	return DeleteProductById(productId)
}

// GetProduct implements ProductRepository
func (*productRepoMock) GetProduct(userId int) ([]*entity.Product, errs.MessageErr) {
	return GetProduct(userId)
}

// GetProductById implements ProductRepository
func (*productRepoMock) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	return GetProductById(productId)
}

// UpdateProductById implements ProductRepository
func (*productRepoMock) UpdateProductById(payload entity.Product) errs.MessageErr {
	return UpdateProductById(payload)
}

func NewProductRepoMock() ProductRepository {
	return &productRepoMock{}
}
