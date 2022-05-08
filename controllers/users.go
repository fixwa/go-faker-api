package controllers

import (
	"github.com/fixwa/go-faker-api/common"
	"github.com/fixwa/go-faker-api/db"
	"github.com/fixwa/go-faker-api/fakery"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var newUser db.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := common.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := db.CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func ListUsers(c *gin.Context) {
	users, restErr := db.ListUsers()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, users)
}

func FindUser(c *gin.Context) {
	userEmail := c.Query("email")
	if userEmail == "" {
		restErr := common.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := db.FindUser(userEmail)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	userEmail := c.Query("email")
	if userEmail == "" {
		restErr := common.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	restErr := db.DeleteUser(userEmail)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}

func UpdateUser(c *gin.Context) {
	userEmail := c.Query("email")
	field := c.Query("field")
	value := c.Query("value")
	if userEmail == "" {
		restErr := common.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if field == "" {
		restErr := common.BadRequest("no field.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if value == "" {
		restErr := common.BadRequest("no value.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := db.UpdateUser(userEmail, field, value)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func GenerateFakeUser(c *gin.Context) {
	fakery.Generator()
	u := db.GenerateFakeUser()
	c.JSON(http.StatusOK, u)
}
