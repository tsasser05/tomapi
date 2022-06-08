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

// Search by first name
func FindNameByFirstName(c *gin.Context) {
	var name models.Name

	if err := models.DB.Where("first_name = ?", c.Param("first_name")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"controllers::names.go::FindNameByFirstName():  Error:": "Record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": name})

}

func UpdateNameByID(c *gin.Context) {
	// Get model
	var name models.Name

	if err := models.DB.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	var input models.UpdateNameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&name).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": name})

}

func DeleteNameByID(c *gin.Context) {
	var name models.Name

	if err := models.DB.Where("id = ?", c.Param("id")).First(&name).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&name)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
