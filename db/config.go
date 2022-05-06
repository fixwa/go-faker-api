package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db               *mongo.Database
	MONGODB_USER     = "root"
	MONGODB_PASSWORD = "root"
)

func ConnectDatabase() {
	if envMongoDbUser := os.Getenv("MONGODB_USER"); envMongoDbUser != "" {
		MONGODB_USER = envMongoDbUser
	}

	if envMongoDbPassword := os.Getenv("MONGODB_PASSWORD"); envMongoDbPassword != "" {
		MONGODB_PASSWORD = envMongoDbPassword
	}

	uri := "mongodb+srv://" + MONGODB_USER + ":" + MONGODB_PASSWORD + "@cluster0.zjyay.mongodb.net/fakerAPI?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	db = client.Database("fakerAPI")
	fmt.Println("Successfuly connected to the database.")
}
