package middleware

import (
	"github.com/Bezunca/API/internal/app/context"
	"github.com/Bezunca/API/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Middleware function setups middleware in echo instance
func Middleware(e *echo.Echo) {
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		context.InjectBezuncAPIContext,
	)
}

// PreMiddleware registers echo's Pre Middleware
func PreMiddleware(e *echo.Echo) {
	if config.Get().ACME.UseTLS {
		e.Pre(middleware.HTTPSRedirect())
	}
}