package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tsasser05/tomapi/models"
)

// Get all names in DB
func GetNames(c *gin.Context) {
	var names []models.Name
	models.DB.Find(&names)
	c.IndentedJSON(http.StatusOK, names)
}

// Create new entry in DB with POST

func CreateName(c *gin.Context) {
	var input models.CreateNameInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	name := models.Name{First_name: input.First_name, Last_name: input.Last_name}
	log.Printf("controllers::CreateName():  name is %v", name)
	models.DB.Create(&name)
	c.JSON(http.StatusOK, gin.H{"data": name})

}

/**************************************************
// Search by first name
func getUserByFirstName(c *gin.Context) {
	fn := c.Param("first_name")

	for _, a := range users {
		if a.First_name == fn {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})

}
**************************************************/
