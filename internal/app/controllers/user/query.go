package user

import (
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByLoginCredentials(ctx echo.Context, loginCredentials models.LoginCredentials) (models.User, error) {

	filter := bson.M{
		"login_credentials": bson.M{
			"email":    loginCredentials.Email,
			"password": loginCredentials.Password,
		},
	}

	user, err := database.GetUser(ctx, filter)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmail(ctx echo.Context, email string) (models.User, error) {

	filter := bson.M{"login_credentials.email": email}
	user, err := database.GetUser(ctx, filter)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
