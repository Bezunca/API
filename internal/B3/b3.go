package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Security struct {
	Ticker string  // Security symbol
	Price  float32 // Security's value 15 min in the past
}

func GetCurrentPrice(ticker string) Security {
	// Gets the price of a given ticker through B3's public API
	// Prices are 15 minutes in the past

	date := fmt.Sprint(time.Now().Format("2006-01-02"))
	url := fmt.Sprintf("https://arquivos.b3.com.br/apinegocios/ticker/%v/%v", ticker, date)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	price, err := jsonparser.GetFloat(responseData, "values", "[0]", "[2]")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return Security{ticker, float32(price)}
}

func GetCurrentPriceList(tickers []string) []Security {
	// Gets the price given a list of security ticker symbols
	prices := make([]Security, len(tickers))
	for i, ticker := range tickers {
		prices[i] = GetCurrentPrice(ticker)
	}
	return prices
}

func main() {
	prices := GetCurrentPriceList([]string{"ITSA4", "FLRY3", "IVVB11"})
	fmt.Println(prices)
}
