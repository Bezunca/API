package middleware

import (
	"github.com/Bezunca/API/internal/config"
	"github.com/Bezunca/API/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		configs := config.Get()

		tokenString := ctx.Request().Header.Get("Authorization")

		user, err := utils.ValidateToken(ctx, tokenString, configs.JWT.SecretAuth)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]map[string]string{"errors": {"general": "Token inválido"}})
		}

		ctx.Set("user", user)

		// TODO: Renovar Token quando estiver perto de expirar

		return next(ctx)
	}
}
