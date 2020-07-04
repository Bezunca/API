package database

import (
	"bezuncapi/internal/models"
	"bezuncapi/internal/parsers"
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUserByEmail(ctx echo.Context, email string) (models.User, error) {

	filter := bson.M{"auth_credentials.email": email}
	findOptions := options.Find()

	usersInterface, err := FindDocuments(ctx, UserDatabase, UsersCollection, filter, findOptions, parsers.ParseUsers)
	if err != nil {
		return models.User{}, err
	}

	users := usersInterface.([]models.User)

	if len(users) == 0 {
		return models.User{}, errors.New("não existe usuário com este email")
	}

	return users[0], nil
}
