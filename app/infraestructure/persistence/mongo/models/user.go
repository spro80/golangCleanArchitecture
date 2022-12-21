package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	IdUser    string             `bson:"idUser,omitempty"`
	Rut       string             `bson:"rut,omitempty"`
	FirstName string             `bson:"firstName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
	Email     string             `bson:"email,omitempty"`
	UserName  string             `bson:"userName,omitempty"`
	Password  string             `bson:"password,omitempty"`
}
