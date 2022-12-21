package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Rut       string             `bson:"rut,omitempty"`
	UserName  string             `bson:"userName,omitempty"`
	Password  string             `bson:"password, omitempty"`
	Email     string             `bson:"email, omitempty"`
	FirstName string             `bson:"firstName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
	Valid     bool               `bson:"valid,omitempty"`
}
