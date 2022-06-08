package models

type Name struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type CreateNameInput struct {
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name" binding:"required"`
}

type UpdateNameInput struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
