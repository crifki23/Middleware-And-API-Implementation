package dto

import "time"

type NewProductRequest struct {
	Title       string `json:"title" valid:"required~title cannot be empty" example:"Jelangkung"`
	Description string `json:"description" valid:"required~description cannot be empty" example:"Cerita fiktif dari Indonesia"`
	Price       int    `json:"price" valid:"required~price cannot be empty" example:"20000"`
}
type NewProductResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
type ProductResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type GetProductResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Data       []ProductResponse `json:"data"`
}
