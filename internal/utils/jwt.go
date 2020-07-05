package utils

import (
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

var AuthExpiration = time.Hour * 24 * 7
var EmailExpiration = time.Hour * 24

func CreateToken(user models.User, expiration int64, secretKey string) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["user_email"] = user.AuthCredentials.Email
	atClaims["expiration"] = expiration

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secretKey))

	return token, err
}

func DecodeToken(tokenString, secretKey string) (jwt.MapClaims, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ValidateToken(ctx echo.Context, token, secret string) (models.User, error) {

	decoded, err := DecodeToken(token, secret)
	if err != nil {
		return models.User{}, err
	}

	if decoded["user_email"] == nil || decoded["expiration"] == nil {
		return models.User{}, errors.New("token inválido")
	}

	userObj, err := database.GetUserByEmail(ctx, decoded["user_email"].(string))
	if err != nil {
		return models.User{}, err
	}

	if float64(time.Now().Unix()) - (decoded["expiration"].(float64)) > 0 {
		return models.User{}, err
	}

	return userObj, nil
}