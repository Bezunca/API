package app

import (
	"bezuncapi/internal/app/controllers/b3"
	"bezuncapi/internal/app/controllers/cei"
	"net/http"

	"github.com/labstack/echo/v4"
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
}

// Placeholder Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
