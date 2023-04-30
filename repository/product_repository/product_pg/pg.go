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
		price = $4
		WHERE id = $1;
	`
	deleteBookByIdQuery = `
		DELETE FROM "products" WHERE id = $1;
	`
	getProductQuery = `
		SELECT id, title, description, price, userId, createdAt, updatedAt from "products"
		WHERE userId = $1;
	`
)

type productPG struct {
	db *sql.DB
}

// DeleteProductById implements product_repository.ProductRepository
func (p *productPG) DeleteProductById(productId int) errs.MessageErr {
	result, err := p.db.Exec(deleteBookByIdQuery, productId)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	if rowsAffected == 0 {
		return errs.NewNotFoundError("product not found")
	}
	return nil
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
func (p *productPG) GetProduct(userId int) ([]*entity.Product, errs.MessageErr) {
	rows, err := p.db.Query(getProductQuery, userId)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	defer rows.Close()
	products := []*entity.Product{}
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.UserId, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}
		products = append(products, &product)
	}
	return products, nil
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
	err := row.Scan(&product.Id, &product.Title, &product.Description, &product.Price, &product.UserId, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("product not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &product, nil
}
