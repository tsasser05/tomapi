package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tsasser05/tomapi/controllers"
	"github.com/tsasser05/tomapi/models"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	router := gin.Default()

	models.ConnectDB()

	router.GET("/", IndexHandler)

	router.GET("/names", controllers.GetNames)
	router.POST("/names", controllers.CreateName)
	router.GET("/names/:first_name", controllers.FindNameByFirstName)
	router.PATCH("/names/:id", controllers.UpdateNameByID)
	router.DELETE("/names/:id", controllers.DeleteNameByID)

	router.Run(":8080")
}
