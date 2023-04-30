package product_repository

import (
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
)

type ProductRepository interface {
	CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById(productId int) (*entity.Product, errs.MessageErr)
	UpdateProductById(payload entity.Product) errs.MessageErr
	DeleteProductById(productId int) errs.MessageErr
	GetProduct(userId int) ([]*entity.Product, errs.MessageErr)
}
