package app

import (
	"bezuncapi/internal/app/controllers/b3"
	"bezuncapi/internal/app/controllers/cei"
	"bezuncapi/internal/app/controllers/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Routes function setups routes in echo instance
func Routes(router *echo.Router) {
	router.Add(http.MethodGet, "/", hello)

	// B3 stuff
	router.Add(http.MethodGet, "/fetch_latest_prices", b3.FetchLatestPrices)

	// CEI stuff
	router.Add(http.MethodGet, "/user_trades", cei.UserTrades)
	router.Add(http.MethodGet, "/user_portfolio_from_trades", cei.UserPortfolioFromTrades)
	router.Add(http.MethodGet, "/user_dividends", cei.UserDividends)
	router.Add(http.MethodGet, "/user_portfolio", cei.UserPortfolio)

	// User
	router.Add(http.MethodPost, "user_login", user.Login)
	router.Add(http.MethodGet, "user_info", UserAuth(user.Info))
}

// Placeholder Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
