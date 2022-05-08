package db

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/fixwa/go-faker-api/common"
	"github.com/fixwa/go-faker-api/fakery"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreatePerson(person *Person) (*Person, *common.RestErr) {
	personsCollection := db.Collection("controllers")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := personsCollection.InsertOne(ctx, bson.M{
		"firstName": person.FirstName,
		"lastName":  person.LastName,
		"gender":    person.Gender,
		"title":     person.Title,
	})
	if err != nil {
		restErr := common.InternalErr("can't insert person to the database.")
		return nil, restErr
	}
	person.ID = result.InsertedID.(primitive.ObjectID)
	return person, nil
}

func ListPersons() (*[]Person, *common.RestErr) {
	personsCollection := db.Collection("controllers")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cur, err := personsCollection.Find(ctx, bson.D{})
	if err != nil {
		restErr := common.InternalErr("Could not return the list of Persons.")
		return nil, restErr
	}

	var results []Person
	if err = cur.All(ctx, &results); err != nil {
		restErr := common.NotFound("Persons not found.")
		return nil, restErr
	}
	return &results, nil
}

func GenerateFakePerson() Person {
	var person Person
	fakery.Generator()
	faker.FakeData(&person)
	return person
}
