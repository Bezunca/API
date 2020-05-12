package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"os"
)

type Security struct{
	Ticker string
	Price float32
}

func GetCurrentPrice(ticker string) Security {
	url := fmt.Sprintf("https://arquivos.b3.com.br/apinegocios/ticker/%v/2020-05-12", ticker)
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
