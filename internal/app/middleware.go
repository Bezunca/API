package app

import (
	"bezuncapi/internal/config"
	"bezuncapi/internal/models"
	"bezuncapi/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Middleware function setups middleware in echo instance
func Middleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func UserAuth(next func(ctx echo.Context, user models.User) error) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		configs := config.Get()

		tokenString := ctx.Request().Header.Get("Authorization")

		userObj, err := utils.ValidateToken(ctx, tokenString, configs.JWT.SecretAuth)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]map[string]string{"errors": {"general": "Token inválido"}})
		}

		// TODO: Renovar Token quando estiver perto de expirar

		return next(ctx, userObj)
	}
}
