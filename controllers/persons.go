package controllers

import (
	"github.com/fixwa/go-faker-api/common"
	"github.com/fixwa/go-faker-api/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePerson(c *gin.Context) {
	var newPerson db.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		restErr := common.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	person, restErr := db.CreatePerson(&newPerson)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, person)
}

func ListPersons(c *gin.Context) {
	persons, restErr := db.ListPersons()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, persons)
}

func GenerateFakePerson(c *gin.Context) {
	person := db.GenerateFakePerson()
	c.JSON(http.StatusOK, person)
}
