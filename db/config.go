package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnectDatabase() {
	uri := "mongodb+srv://superbisor:MH3GuRIb096HdRzv@cluster0.zjyay.mongodb.net/fakerAPI?retryWrites=true&w=majority"
	//uri := "mongodb://root:root@localhost:27017"
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
