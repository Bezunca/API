package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func createToken(user User) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id

	//TODO: Token expiration

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("3K2jwcqZEQP5hnogXu0j")))

	return token, err
}

func Login(ctx echo.Context) error {

	userEmail, userPassword, err := validateUserLogin(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]error{"error": err})
	}

	//TODO: Check user (MongoDB)
	user := User{
		Id: 999,
		Email: userEmail,
		Password: userPassword,
	}
	token, err := createToken(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]error{"error": err})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func Info(ctx echo.Context, user User) error {
	return ctx.JSON(http.StatusOK, map[string]int{"id": user.Id})
}