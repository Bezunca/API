package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func validateUserLogin(ctx echo.Context) (string, string, error) {
	userEmail, userPassword, ok := ctx.Request().BasicAuth()
	if !ok {
		return "", "", echo.NewHTTPError(http.StatusBadRequest, "Missing user's e-mail and password in HTTP Basic Auth")
	}
	//TODO: Validate email and password
	return userEmail, userPassword, nil
}
