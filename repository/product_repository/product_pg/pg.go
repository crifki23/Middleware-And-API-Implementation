package product_pg

import (
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/repository/product_repository"
	"database/sql"
	"errors"
	"fmt"
)

const (
	getProductByIdQuery = `
		SELECT id, title, description, price, userId, createdAt, updatedAt from "products"
		WHERE id = $1;
	`

	updateProductByIdQuery = `
		UPDATE "products"
		SET title = $2,
		description = $3,
		price = $4,
		updateAt = $5
		WHERE id = $1;
	`
)

type productPG struct {
	db *sql.DB
}

// UpdateProductById implements product_repository.ProductRepository
func (p *productPG) UpdateProductById(payload entity.Product) errs.MessageErr {
	_, err := p.db.Exec(updateProductByIdQuery, payload.Id, payload.Title, payload.Description, payload.Price)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

func NewProductPG(db *sql.DB) product_repository.ProductRepository {
	return &productPG{
		db: db,
	}
}

// GetMovies implements product_repository.ProductRepository
func (*productPG) GetMovies() ([]*entity.Product, errs.MessageErr) {
	return nil, nil
}

func (p *productPG) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	createProductQuery := `
		INSERT INTO "products"
		(
			title,
			description,
			price,
			userId
		)
		VALUES($1, $2, $3, $4)
		RETURNING id,title, description, price, userId;
	`
	row := p.db.QueryRow(createProductQuery, productPayload.Title, productPayload.Description, productPayload.Price, productPayload.UserId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}
func (p *productPG) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	row := p.db.QueryRow(getProductByIdQuery, productId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.UserId, &product.Description, &product.Price, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("product not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}
