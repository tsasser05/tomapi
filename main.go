package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tsasser05/tomapi/controllers"
	"github.com/tsasser05/tomapi/models"
)

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.GET("/names", controllers.GetNames)
	//router.POST("/names", postUsers)
	//router.GET("/names/:first_name", getUserByFirstName)

	router.Run("localhost:8080")
}
