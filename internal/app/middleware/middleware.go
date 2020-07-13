package middleware

import (
	"bezuncapi/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Middleware function setups middleware in echo instance
func Middleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

// PreMiddleware registers echo's Pre Middleware
func PreMiddleware(e *echo.Echo) {
	if config.Get().ACME.UseTLS {
		e.Pre(middleware.HTTPSRedirect())
	}
}