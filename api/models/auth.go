package models

import (
	"time"
)

type User struct {
	Id        string     `json:"id"`
	FullName  string     `json:"full_name"`
	Login     string     `json:"login"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	UserType  string     `json:"user_type"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CreateUserModel struct {
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

// type GetAllCategoryModel struct {
// 	Categories []Category `json:"products"`
// 	Count     int32           `json:"count"`
// }
// type UpdateCategoryModel struct {
// 	Title string `json:"title" binding:"required"`
// }
