package models

type Name struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
