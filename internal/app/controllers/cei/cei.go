package cei

import (
	"bezuncapi/internal/app/controllers/cei/scraper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func extractUserCPFAndCEIPassword(ctx echo.Context) (string, string, error) {
	userCPF, ceiPassword, ok := ctx.Request().BasicAuth()
	if !ok {
		return "", "", echo.NewHTTPError(http.StatusBadRequest, "Missing user's CPF and CEI password in HTTP Basic Auth")

	}
	return userCPF, ceiPassword, nil
}

func UserTrades(ctx echo.Context) error {
	userCPF, ceiPassword, err := extractUserCPFAndCEIPassword(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, scraper.GetUserTrades(userCPF, ceiPassword))
}

func UserPortfolioFromTrades(ctx echo.Context) error {
	userCPF, ceiPassword, err := extractUserCPFAndCEIPassword(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, scraper.GetUserPortfolioFromTrades(userCPF, ceiPassword))
}

func UserDividends(ctx echo.Context) error {
	userCPF, ceiPassword, err := extractUserCPFAndCEIPassword(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, scraper.GetUserDividends(userCPF, ceiPassword))
}

func UserPortfolio(ctx echo.Context) error {
	userCPF, ceiPassword, err := extractUserCPFAndCEIPassword(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, scraper.GetUserPortfolio(userCPF, ceiPassword))
}
