package models

import "time"

type Product struct{
     Id string `json:"id"`   
     Title string `json:"title" binding:"required"`
     Description string `json:"description" binding:"required"`
     Quantity int `json:"quantity" binding:"required"`
     Price int `json:"price" binding:"required"`
     Category_id string `json:"category_id" binding:"required"`
     CreatedAt time.Time  `json:"created_at"`
	 UpdatedAt *time.Time `json:"updated_at"`
}

type CreateProductModel struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity int32 `json:"quantity" binding:"required"`
	Price int32 `json:"price" binding:"required"`
	Category_id string `json:"category_id" binding:"required"`
}

type UpdateProductModel struct {
	Id string `json:"id" binding:"required"`   
	Title string `json:"title"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
}

type GetProductByIdModel struct {
	Id string `json:"id"`   
	Title string `json:"title"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
	Category Category `json:"category"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
type GetAllProductModel struct {
	Products []Product `json:"products"`
	Count     int32           `json:"count"`
}

