package app

import (
	"context"
	internalContext "github.com/Bezunca/API/internal/app/context"
	"github.com/Bezunca/API/internal/app/controllers/auth"
	"github.com/Bezunca/API/internal/app/controllers/b3"
	"github.com/Bezunca/API/internal/app/controllers/wallet"
	"github.com/Bezunca/API/internal/app/middleware"
	"github.com/Bezunca/API/internal/graph"
	"github.com/Bezunca/API/internal/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/labstack/echo/v4"
)

// Routes function setups routes in echo instance
func Routes(echo *echo.Echo) {
	echo.GET( "/", hello)

	// B3 stuff
	echo.GET( "/fetch_latest_prices", b3.FetchLatestPrices)

	// Auth
	echo.POST( "/auth/register", auth.Register)
	echo.POST( "/auth/confirm_registration", auth.ConfirmRegistration)
	echo.POST( "/auth/forgot_password", auth.ForgotPassword)
	echo.POST( "/auth/reset_password", auth.ResetPassword)
	echo.POST( "/auth/login", auth.Login)

	loggedRoutes := echo.Group("", middleware.UserAuth)
	loggedRoutes.GET( "/auth/info", auth.Info)
	loggedRoutes.POST( "/wallet/cei_sync", wallet.CEISync)
	loggedRoutes.POST( "/query", graphqlHandler)
}

// Placeholder Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

// GraphQL Handler
func graphqlHandler(c echo.Context) error {
	ctx := c.(*internalContext.BezuncAPIContext)
	user := ctx.User()

	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	h.ServeHTTP(
		c.Response(),
		ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), "user", user)),
	)
	return nil
}
