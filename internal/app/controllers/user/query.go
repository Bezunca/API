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

func PostUser(ctx echo.Context, user models.User) bool {

	users := make([]interface{}, 1)
	users[0] = user

	inserted := database.InsertDocuments(
		ctx,
		database.UserDatabase,
		database.UsersCollection,
		users,
	)

	return inserted
}

func GetUserByEmail(ctx echo.Context, email string) (models.User, error) {

	filter := bson.M{"auth_credentials.email": email}
	findOptions := options.Find()

	usersInterface, err := database.FindDocuments(ctx, database.UserDatabase, database.UsersCollection, filter, findOptions, parsers.ParseUsers)
	if err != nil {
		return models.User{}, err
	}

	users := usersInterface.([]models.User)

	if len(users) == 0 {
		return models.User{}, errors.New("não existe usuário com este email")
	}

	return users[0], nil
}

func UpdateUserRegisterConfirmation(ctx echo.Context, email string) bool {

	filter := bson.M{"auth_credentials.email": email}
	update := bson.D{
		{"$set", bson.D{
				{"auth_credentials.activated", true},
			},
		},
	}

	updated := database.UpdateDocuments(ctx, database.UserDatabase, database.UsersCollection, filter, update)
	return updated
}

func UpdateUserResetPassword(ctx echo.Context, email, password string) bool {

	filter := bson.M{"auth_credentials.email": email}
	update := bson.D{
		{"$set", bson.D{
			{"auth_credentials.activated", true},
			{"auth_credentials.password", password},
		},
		},
	}

	updated := database.UpdateDocuments(ctx, database.UserDatabase, database.UsersCollection, filter, update)
	return updated
}
