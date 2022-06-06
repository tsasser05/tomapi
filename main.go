package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

var users = []user{
	{First_name: "Foo", Last_name: "Bar"},
	{First_name: "Baz", Last_name: "Baz"},
	{First_name: "Moo", Last_name: "Cow"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUsers)
	router.GET("/users/:first_name", getUserByFirstName)
	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

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
