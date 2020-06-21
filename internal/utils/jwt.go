package utils

import (
	"bezuncapi/internal/models"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user models.User, secretKey string) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["user_email"] = user.AuthCredentials.Email

	//TODO: Token expiration

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
