package validators

import (
	"bezuncapi/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ValidateUserLogin(ctx echo.Context) (models.AuthCredentials, error) {

	userEmail, userPassword, ok := ctx.Request().BasicAuth()
	if !ok {
		return models.AuthCredentials{}, echo.NewHTTPError(http.StatusBadRequest, "Missing user's e-mail and password in HTTP Basic Auth")
	}

	authCredentials := models.AuthCredentials{
		Email:    userEmail,
		Password: userPassword,
	}

	err := ValidateStruct(authCredentials)
	if err != nil {
		return models.AuthCredentials{}, err
	}

	return authCredentials, nil
}

func ValidateUserRegister(ctx echo.Context) (models.User, error) {

	user := models.User{}
	if err := ctx.Bind(&user); err != nil {
		return models.User{}, err
	}

	err := ValidateStruct(user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
