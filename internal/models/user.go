package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name               string             `bson:"name" json:"name"`
	AuthCredentials    AuthCredentials    `bson:"auth_credentials" json:"auth_credentials"`
	WalletsCredentials WalletsCredentials `bson:"wallets_credentials" json:"wallets_credentials"`
}

type AuthCredentials struct {
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
	Activated bool   `bson:"activated" json:"activated"`
}
