package context

import (
	"context"
	"github.com/Bezunca/API/internal/models"
	"github.com/labstack/echo/v4"
)

// DockingBayContext envelops echo.Context with additional functionality
type BezuncAPIContext struct {
	echo.Context
}

// User returns user struct for current request
func (ctx *BezuncAPIContext) User() models.User {
	return ctx.Get("user").(models.User)
}

// DockingBayContext envelops echo.Context with additional functionality
type GraphQLContext struct {
	context.Context
}

// User returns user struct for current request
func (ctx *GraphQLContext) User() models.User {
	return ctx.Value("user").(models.User)
}

// InjectBezuncAPIContext is a function to be used as a middleware to envelope DockingBayContext on echo.Context
func InjectBezuncAPIContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return next(&BezuncAPIContext{ctx})
	}
}
