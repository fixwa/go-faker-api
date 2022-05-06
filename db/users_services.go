package db

import (
	"context"
	"github.com/fixwa/go-faker-api/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateUser(user *User) (*User, *common.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		restErr := common.InternalErr("can't insert user to the database.")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

func ListUsers() (*[]User, *common.RestErr) {
	usersCollection := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cur, err := usersCollection.Find(ctx, bson.D{})
	if err != nil {
		restErr := common.InternalErr("Could not return the list of Persons.")
		return nil, restErr
	}

	var results []User
	
	if err = cur.All(ctx, &results); err != nil {
		restErr := common.NotFound("Users not found.")
		return nil, restErr
	}
	return &results, nil
}

func FindUser(email string) (*User, *common.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		restErr := common.NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}

func DeleteUser(email string) *common.RestErr {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		restErr := common.NotFound("faild to delete.")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := common.NotFound("user not found.")
		return restErr
	}
	return nil
}

func UpdateUser(email string, field string, value string) (*User, *common.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := common.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := common.NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := common.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := FindUser(email)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
