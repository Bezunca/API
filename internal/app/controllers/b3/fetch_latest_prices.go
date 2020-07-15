package b3

import (
	"github.com/Bezunca/API/internal/config"
	"net/http"

	"github.com/Bezunca/b3lib/fetch_price"
	"github.com/labstack/echo/v4"
)

var fetcher func(tickers []string) ([]fetch_price.FetchedPrice, []error)

func SetFetcher() {
	fetcher = fetch_price.New(config.Get().B3CacheTimeout, &http.Client{})
}

// FetchLatestPrices gets the latest prices of securities through B3's public API
// Prices are 15 minutes in the past
func FetchLatestPrices(ctx echo.Context) error {
	tickers, ok := ctx.QueryParams()["tickers"]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing tickers query parameter")
	}

	prices, err := fetcher(tickers)
	if err != nil {
		ctx.Logger().Error(err)
		return echo.NewHTTPError(http.StatusServiceUnavailable, "Unable to fetch prices from B3, try again later")
	}

	return ctx.JSON(http.StatusOK, prices)
}
