package app

import (
	"bezuncapi/internal/app/controllers/b3"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(router *echo.Router) {
	router.Add("GET", "/", hello)
	router.Add("GET", "/fetch_quotation", b3.FetchQuotation)
}

// Placeholder Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
