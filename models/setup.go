package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("sqlite3", "names.db")

	if err != nil {
		panic("Failed to connect to names database")
	}

	database.AutoMigrate(&Name{})
	DB = database
}
