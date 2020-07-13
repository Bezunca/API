package context

import (
	"bezuncapi/internal/models"
	"context"
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
