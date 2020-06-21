package user

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"bezuncapi/internal/validators"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

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

	err = sendRegisterEmail(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func ConfirmRegistration(ctx echo.Context) error {

	tokenString := strings.Split(ctx.Request().URL.Path, "/")[3]

	configs := config.Get()
	decoded, err := utils.DecodeToken(tokenString, configs.JWTSecretEmail)
	if err != nil || decoded["user_email"] == nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
	}

	user, err := GetUserByEmail(ctx, decoded["user_email"].(string))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
	}

	updated := UpdateUserRegisterConfirmation(ctx, user.AuthCredentials.Email)
	if !updated {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
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
	fmt.Println(user)
	if !user.AuthCredentials.Activated {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "usuário não está ativado"})
	}

	configs := config.Get()
	token, err := utils.CreateToken(user, configs.JWTSecretAuth)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]error{"error": err})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}

func Info(ctx echo.Context, user models.User) error {
	return ctx.JSON(http.StatusOK, user)
}
