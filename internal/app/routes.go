package app

import (
	"bezuncapi/internal/app/controllers/auth"
	"bezuncapi/internal/app/controllers/b3"
	"bezuncapi/internal/app/controllers/wallet"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Routes function setups routes in echo instance
func Routes(router *echo.Router) {
	router.Add(http.MethodGet, "/", hello)

	// B3 stuff
	router.Add(http.MethodGet, "/fetch_latest_prices", b3.FetchLatestPrices)

	// Auth
	router.Add(http.MethodPost, "/auth/register", auth.Register)
	router.Add(http.MethodPost, "/auth/confirm_registration", auth.ConfirmRegistration)
	router.Add(http.MethodPost, "/auth/forgot_password", auth.ForgotPassword)
	router.Add(http.MethodPost, "/auth/reset_password", auth.ResetPassword)
	router.Add(http.MethodPost, "/auth/login", auth.Login)
	router.Add(http.MethodGet, "/auth/info", UserAuth(auth.Info))

	// Wallets
	router.Add(http.MethodPost, "/wallet/cei_sync", UserAuth(wallet.CEISync))
}

// Placeholder Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
