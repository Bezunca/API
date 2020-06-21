package user

import (
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"
	"bezuncapi/internal/parsers"
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(ctx echo.Context, filter bson.M, findOptions *options.FindOptions) ([]models.User, error) {

	users, err := database.FindDocuments(
		ctx,
		database.UserDatabase,
		database.UsersCollection,
		filter,
		findOptions,
		parsers.ParseUsers,
	)
	if err != nil {
		return []models.User{}, err
	}

	return users.([]models.User), nil
}

func PostUser(ctx echo.Context, user models.User) bool {

	users := make([]interface{}, 1)
	users[0] = user

	inserted := database.InsertDocuments(
		ctx,
		database.UserDatabase,
		database.UsersCollection,
		users)

	return inserted
}

func GetUserByEmail(ctx echo.Context, email string) (models.User, error) {

	filter := bson.M{"auth_credentials.email": email}
	findOptions := options.Find()

	users, err := GetUsers(ctx, filter, findOptions)

	if err != nil {
		return models.User{}, err
	}

	if len(users) == 0 {
		return models.User{}, errors.New("não existe usuário com este email")
	}

	return users[0], nil
}
