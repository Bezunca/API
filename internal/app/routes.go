package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(router *echo.Router) {
	router.Add("GET", "/", hello)
}


// Placeholder Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
