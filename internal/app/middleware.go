package app

import (
	"bezuncapi/internal/app/controllers/user"
	"bezuncapi/internal/config"
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

func injectMongoClient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		mongoClient, err := database.GetConnection()
		if err != nil {
			log.Fatal(err)
		}
		ctx.Set("mongoClient", mongoClient)
		return next(ctx)
	}
}

// Middleware function setups middleware in echo instance
func Middleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(injectMongoClient)
}

func UserAuth(next func(ctx echo.Context, user models.User) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		configs := config.Get()

		tokenString := ctx.Request().Header.Get("Authorization")

		decoded, err := utils.DecodeToken(tokenString, configs.JWTSecretAuth)
		if err != nil || decoded["user_email"] == nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
		}

		userObj, err := user.GetUserByEmail(ctx, decoded["user_email"].(string))
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Token"})
		}

		return next(ctx, userObj)
	}
}
