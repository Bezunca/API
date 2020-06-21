package user

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"bezuncapi/internal/validators"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func createToken(user models.User) (string, error) {

	configs := config.Get()

	atClaims := jwt.MapClaims{}
	atClaims["user_email"] = user.AuthCredentials.Email

	//TODO: Token expiration

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(configs.JWTSecret))

	return token, err
}

func Register(ctx echo.Context) error {

	user, err := validators.ValidateUserRegister(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	_, err = GetUserByEmail(ctx, user.AuthCredentials.Email)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "já existe uma conta cadastrada com esse email"})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.AuthCredentials.Password), bcrypt.MinCost)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]error{"error": err})
	}

	user.AuthCredentials.Password = string(hashPassword)

	inserted := PostUser(ctx, user)
	if !inserted {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "usuário não foi inserido no banco"})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func Login(ctx echo.Context) error {

	authCredentials, err := validators.ValidateUserLogin(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := GetUserByEmail(ctx, authCredentials.Email)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.AuthCredentials.Password), []byte(authCredentials.Password))
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
