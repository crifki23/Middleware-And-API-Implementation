package entity

import (
	"chapter3-sesi2/dto"
	"time"
)

type Product struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	UserId      int       `json:"userId"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) EntityToProductResponseDto() dto.ProductResponse {
	return dto.ProductResponse{
		Id:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
