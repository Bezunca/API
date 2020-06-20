package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID
	LoginCredentials LoginCredentials
}

type LoginCredentials struct {
	Email    string
	Password string
}
