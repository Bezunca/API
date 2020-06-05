package b3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type b3QuotationColumns struct {
	Name            string `json:"name"`
	FriendlyName    string `json:"friendlyName"`
	FriendlyNamePt  string `json:"friendlyNamePt"`
	FriendlyNameEn  string `json:"friendlyNameEn"`
	Type            int    `json:"type"`
	Format          string `json:"format"`
	ColumnAlignment int    `json:"columnAlignment"`
	ValueAlignment  int    `json:"valueAlignment"`
}

type b3QuotationResponse struct {
	Name         string               `json:"name"`
	FriendlyName string               `json:"friendlyName"`
	Columns      []b3QuotationColumns `json:"columns"`
	Values       [][6]interface{}     `json:"values"`
	// Inner slices follows columns order
}

type fetchedQuotation struct {
	Ticker string
	Price  float64
}

// FetchQuotation gets the price of a given tickers through B3's public API
// Prices are 15 minutes in the past
func FetchQuotation(ctx echo.Context) error {
	tickers, ok := ctx.QueryParams()["tickers"]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing tickers query parameter")
	}

	today := time.Now()
	if today.Hour() < 10 {
		// B3 only got quotations values after market opening, so before 10AM, we get yesterday values
		today = today.Add(-24 * time.Hour)
	}
	date := today.Format("2006-01-02")

	fetchedQuotationsChan := make(chan fetchedQuotation, len(tickers))
	errChan := make(chan error)
	defer func() {
		close(fetchedQuotationsChan)
		close(errChan)
	}()

	for _, ticker := range tickers {
		go getCurrentPrice(date, ticker, fetchedQuotationsChan, errChan)
	}

	var fetchedQuotations []fetchedQuotation
	var errors []error
	for i := 0; i < len(tickers); i++ {
		select {
		case fetchedQuotation := <-fetchedQuotationsChan:
			fetchedQuotations = append(fetchedQuotations, fetchedQuotation)
		case err := <-errChan:
			ctx.Logger().Error(err)
			errors = append(errors, err)
		}
	}

	if len(errors) != 0 {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "One or more tickers couldn't be fetched from B3, try again later")
	}

	return ctx.JSON(http.StatusOK, fetchedQuotations)
}

func getCurrentPrice(date string, ticker string, quotations chan<- fetchedQuotation, errChan chan<- error) {
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

	b3Response := new(b3QuotationResponse)
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(b3Response); err != nil {
		errChan <- fmt.Errorf("failed to decode B3 quotation for %v on %v: %v", ticker, date, err)
		return
	}

	if len(b3Response.Values) == 0 {
		errChan <- fmt.Errorf("failed to get B3 quotation values field for %v on %v", ticker, date)
		return
	}

	quotations <- fetchedQuotation{
		Ticker: b3Response.Values[0][0].(string),
		Price:  b3Response.Values[0][2].(float64),
	}
}
