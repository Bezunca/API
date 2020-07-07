package auth

import (
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

func UpdateUserRegisterConfirmation(ctx echo.Context, email string) bool {

	filter := bson.M{"auth_credentials.email": email}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "auth_credentials.activated", Value: true},
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
