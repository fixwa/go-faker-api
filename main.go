package main

import (
	"github.com/fixwa/go-faker-api/controllers"
	"github.com/fixwa/go-faker-api/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	ApiVersion  = "/v1"
	defaultPort = "8080"
)

func main() {
	engine := gin.Default()
	db.ConnectDatabase()

	engine.GET(ApiVersion+"/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"version": 1.0,
		})
	})

	engine.POST(ApiVersion+"/users/create", controllers.CreateUser)
	engine.GET(ApiVersion+"/users/list", controllers.ListUsers)
	engine.GET(ApiVersion+"/users/find", controllers.FindUser)
	engine.GET(ApiVersion+"users/delete", controllers.DeleteUser)
	engine.GET(ApiVersion+"/users/update", controllers.UpdateUser)

	engine.POST(ApiVersion+"/persons/create", controllers.CreatePerson)
	engine.GET(ApiVersion+"/persons/list", controllers.ListPersons)

	if envPort := os.Getenv("PORT"); envPort != "" {
		defaultPort = envPort
	}
	engine.Run(":" + defaultPort)
}
