package main

import (
	"github.com/fixwa/go-faker-api/controllers"
	"github.com/fixwa/go-faker-api/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
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

	engine.GET(ApiVersion+"/fake", controllers.GenerateAllFaked)

	engine.POST(ApiVersion+"/users/create", controllers.CreateUser)
	engine.GET(ApiVersion+"/users/list", controllers.ListUsers)
	engine.GET(ApiVersion+"/users/find", controllers.FindUser)
	engine.GET(ApiVersion+"/users/delete", controllers.DeleteUser)
	engine.GET(ApiVersion+"/users/update", controllers.UpdateUser)
	engine.GET(ApiVersion+"/users/fake", controllers.GenerateFakeUser)

	engine.POST(ApiVersion+"/persons/create", controllers.CreatePerson)
	engine.GET(ApiVersion+"/persons/list", controllers.ListPersons)
	engine.GET(ApiVersion+"/persons/fake", controllers.GenerateFakePerson)

	if envPort := os.Getenv("PORT"); envPort != "" {
		defaultPort = envPort
	}

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Run(":" + defaultPort)
}
