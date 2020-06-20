package app

import (
	"bezuncapi/internal/app/controllers/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// Middleware function setups middleware in echo instance
func Middleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func UserAuth (next func(ctx echo.Context, user user.User) error) echo.HandlerFunc {
	return func(ctx echo.Context) error{

		tokenString := ctx.Request().Header.Get("Authorization")

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("3K2jwcqZEQP5hnogXu0j"), nil
		})
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
		}

		if claims["user_id"] != nil {

			//TODO: Return user (MongoDB)
			user := user.User{
				Id: 999,
				Email: "email",
				Password: "password",
			}

			return next(ctx, user)
		}else {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
		}
	}
}