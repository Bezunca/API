package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Scraping struct {
	ID                 primitive.ObjectID `json:"user_id" bson:"user_id"`
	WalletsCredentials WalletsCredentials `json:"wallets_credentials" bson:"wallets_credentials"`
}
