package user

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

func createToken(user models.User) (string, error) {

	configs := config.Get()

	atClaims := jwt.MapClaims{}
	atClaims["user_email"] = user.LoginCredentials.Email

	//TODO: Token expiration

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(configs.JWTSecret))

	return token, err
}

func Login(ctx echo.Context) error {

	userEmail, userPassword, err := validateUserLogin(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]error{"error": err})
	}

	user, err := GetUserByLoginCredentials(ctx, models.LoginCredentials{
		Email:    userEmail,
		Password: userPassword},
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	token, err := createToken(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]error{"error": err})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func Info(ctx echo.Context, user models.User) error {
	return ctx.JSON(http.StatusOK, user)
}
