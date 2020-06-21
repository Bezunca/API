package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name" json:"name" validate:"required,min=3,max=25"`
	AuthCredentials AuthCredentials    `bson:"auth_credentials" json:"auth_credentials" validate:"required"`
}

type AuthCredentials struct {
	Email     string `bson:"email" json:"email" validate:"required,email"`
	Password  string `bson:"password" json:"password" validate:"required,min=3,max=25"`
	Activated bool   `bson:"activated" json:"activated"`
}
