package b3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type b3PriceColumns struct {
	Name            string `json:"name"`
	FriendlyName    string `json:"friendlyName"`
	FriendlyNamePt  string `json:"friendlyNamePt"`
	FriendlyNameEn  string `json:"friendlyNameEn"`
	Type            int    `json:"type"`
	Format          string `json:"format"`
	ColumnAlignment int    `json:"columnAlignment"`
	ValueAlignment  int    `json:"valueAlignment"`
}

type b3PriceResponse struct {
	Name         string           `json:"name"`
	FriendlyName string           `json:"friendlyName"`
	Columns      []b3PriceColumns `json:"columns"`
	Values       [][6]interface{} `json:"values"`
	// Inner slices follows columns order
}

type Security struct {
	Ticker string
	Price  float64
}

// FetchLatestPrice gets the latest prices of securities through B3's public API
// Prices are 15 minutes in the past
func FetchLatestPrice(ctx echo.Context) error {
	tickers, ok := ctx.QueryParams()["tickers"]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing tickers query parameter")
	}

	today := time.Now()
	todayWeekday := today.Weekday()
	if todayWeekday == time.Sunday {
		today = today.Add(-48 * time.Hour)
	} else if todayWeekday == time.Saturday || today.Hour() < 6 {
		// B3 only provides price values after market opening, so before 6AM, we get yesterday values
		// If it's after 6AM, we zero the prices, waiting for the market to open
		today = today.Add(-24 * time.Hour)
	}
	date := today.Format("2006-01-02")

	fetchedPricesChan := make(chan Security, len(tickers))
	errChan := make(chan error)
	defer func() {
		close(fetchedPricesChan)
		close(errChan)
	}()

	for _, ticker := range tickers {
		go getCurrentPrice(date, ticker, fetchedPricesChan, errChan)
	}

	var fetchedPrices []Security
	var errors []error
	for i := 0; i < len(tickers); i++ {
		select {
		case fetchedPrice := <-fetchedPricesChan:
			fetchedPrices = append(fetchedPrices, fetchedPrice)
		case err := <-errChan:
			ctx.Logger().Error(err)
			errors = append(errors, err)
		}
	}

	if len(errors) != 0 {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "One or more tickers couldn't be fetched from B3, try again later")
	}

	return ctx.JSON(http.StatusOK, fetchedPrices)
}

func getCurrentPrice(date string, ticker string, pricesChannel chan<- Security, errChan chan<- error) {
	url := fmt.Sprintf("https://arquivos.b3.com.br/apinegocios/ticker/%v/%v", ticker, date)
	response, err := http.Get(url)
	if err != nil {
		errChan <- fmt.Errorf("failed to fetch B3 quotation for %v on %v: %v", ticker, date, err)
		return
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			errChan <- fmt.Errorf("failed to close B3 response body for %v on %v: %v", ticker, date, err)
		}
	}()

	b3Response := new(b3PriceResponse)
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(b3Response); err != nil {
		errChan <- fmt.Errorf("failed to decode B3 quotation for %v on %v: %v", ticker, date, err)
		return
	}

	price := 0.0
	if len(b3Response.Values) != 0 {
		price = b3Response.Values[0][2].(float64)
	}

	pricesChannel <- Security{
		Ticker: b3Response.Name,
		Price:  price,
	}
}
