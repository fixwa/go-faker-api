package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty" faker:"customIdFaker"`
	Name     string             `json:"name" bson:"name,omitempty" faker:"first_name"`
	Email    string             `json:"email" bson:"email,omitempty" faker:"email"`
	Password string             `json:"password" bson:"password,omitempty" faker:"password"`
}

type Person struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty" faker:"customIdFaker"`
	FirstName string             `json:"firstName" bson:"firstName,omitempty" faker:"first_name"`
	LastName  string             `json:"lastName" bson:"lastName,omitempty" faker:"last_name"`
	Gender    string             `json:"gender" bson:"gender,omitempty" faker:"gender"`
	Title     string             `json:"title" bson:"title,omitempty" faker:"title_female"`
}
