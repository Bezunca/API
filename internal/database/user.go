package database

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"bezuncapi/internal/parsers"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const UsersCollection = "users"

func GetUser(ctx echo.Context, filter bson.M) (models.User, error) {

	mongoClient := ctx.Get("mongoClient").(*mongo.Client)

	configs := config.Get()
	usersCollection := mongoClient.Database(configs.MongoDatabase).Collection(UsersCollection)

	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var data map[string]interface{}
	err := usersCollection.FindOne(queryCtx, filter).Decode(&data)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	user, ok := parsers.ParseUser(data)
	if !ok {
		return models.User{}, errors.New("cant decode data for the user")
	}

	return user, nil
}
