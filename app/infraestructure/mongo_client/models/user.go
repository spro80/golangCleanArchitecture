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
	Profile   Profile            `bson:"profile,omitempty"`
}

type Profile struct {
	ProfileId       int    `bson:"profileId,omitempty"`
	ProfileStatus   bool   `bson:"profileStatus,omitempty"`
	ProfileDateInit string `bson:"profileDateInit,omitempty"`
	ProfileDateEnd  string `bson:"profileDateEnd,omitempty"`
	ProfileAllTime  bool   `bson:"profileAllTime,omitempty"`
}
