package controllers

import (
	"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/tsasser05/tomapi/models"
)

func GetNames(c *gin.Context) {
	var names []models.Name
	models.DB.Find(&names)
	c.IndentedJSON(http.StatusOK, names)
}

/**************************************************
// POST users

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

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
