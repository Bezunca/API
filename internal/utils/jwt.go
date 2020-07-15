package utils

import (
	"errors"
	"github.com/Bezunca/API/internal/database"
	"github.com/Bezunca/API/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

var AuthExpiration = time.Hour * 24 * 7
var EmailExpiration = time.Hour * 24

func CreateToken(user models.User, expiration int64, secretKey string) (string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = (user.ID).Hex()
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
	bearerToken := strings.Replace(token, "Bearer ", "", 1)
	decoded, err := DecodeToken(bearerToken, secret)
	if err != nil {
		return models.User{}, err
	}

	if decoded["user_id"] == nil || decoded["expiration"] == nil {
		return models.User{}, errors.New("invalid token")
	}

	userId, err := primitive.ObjectIDFromHex(decoded["user_id"].(string))
	if err != nil {
		return models.User{}, errors.New("invalid token")
	}

	userObj, err := database.GetUserByID(ctx, userId)
	if err != nil {
		return models.User{}, err
	}

	if float64(time.Now().Unix())-(decoded["expiration"].(float64)) > 0 {
		return models.User{}, errors.New("expired token")
	}

	return userObj, nil
}
