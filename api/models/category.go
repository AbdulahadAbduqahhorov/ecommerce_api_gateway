package models

import "time"

type Category struct {
	Id        string     `json:"id"`
	Title string `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateCategoryModel struct {
	Title string `json:"title"`
}

type UpdateCategoryModel struct {
	Title string `json:"title" binding:"required"`
}
