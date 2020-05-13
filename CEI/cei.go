package main

import (
	"./base"
)

type Asset struct {
	Symbol       string
	Description  string
	Market       string
	Amount       int
	AveragePrice float32
}

func getUserTrades(cpf, password string) []base.Trade {
	return base.GetRawUserTrades(cpf, password)
}

func getUserPortfolio(cpf, password string) []Asset {
	trades := base.GetRawUserTrades(cpf, password)

	if trades != nil {

		portfolioMap := map[string]*Asset{}

		for _, trade := range trades {

			if _, ok := portfolioMap[trade.Symbol]; !ok {
				portfolioMap[trade.Symbol] = &Asset{
					trade.Symbol,
					trade.Description,
					trade.Market,
					0,
					0,
				}
			}

			if trade.Action == "C" {
				portfolioMap[trade.Symbol].Amount += trade.Amount
				portfolioMap[trade.Symbol].AveragePrice += trade.FullPrice
			} else {
				portfolioMap[trade.Symbol].Amount -= trade.Amount
				portfolioMap[trade.Symbol].AveragePrice -= trade.FullPrice
			}
		}

		var portfolio []Asset

		for _, asset := range portfolioMap {
			asset.AveragePrice = asset.AveragePrice / float32(asset.Amount)
			portfolio = append(portfolio, *asset)
		}
		return portfolio
	} else {
		return nil
	}
}

func main() {

	/*trades := getUserTrades("CPF", "SENHA")
	tradesJson, _ := json.Marshal(trades)
	fmt.Println("\n\n\n", string(tradesJson))*/

	/*portfolio := getUserPortfolio("CPF", "SENHA")
	portfolioJson, _ := json.Marshal(portfolio)
	fmt.Println("\n\n\n", string(portfolioJson))*/
}
