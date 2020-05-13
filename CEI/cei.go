package main

import (
	"./base"
	"github.com/shopspring/decimal"
)

type Asset struct {
	Symbol       string
	Description  string
	Market       string
	Amount       decimal.Decimal
	AveragePrice decimal.Decimal
}

func getUserTrades(cpf, password string) []base.Trade {
	return base.GetRawUserTrades(cpf, password)
}

func getUserPortfolio(cpf, password string) []Asset {
	trades := base.GetRawUserTrades(cpf, password)

	if len(trades) > 0 {

		portfolioMap := map[string]*Asset{}

		for _, trade := range trades {

			if _, ok := portfolioMap[trade.Symbol]; !ok {
				portfolioMap[trade.Symbol] = &Asset{
					trade.Symbol,
					trade.Description,
					trade.Market,
					decimal.NewFromInt(0),
					decimal.NewFromInt(0),
				}
			}

			if trade.Action == "C" {
				portfolioMap[trade.Symbol].Amount = portfolioMap[trade.Symbol].Amount.Add(trade.Amount)
				portfolioMap[trade.Symbol].AveragePrice = portfolioMap[trade.Symbol].AveragePrice.Add(trade.FullPrice)
			} else {
				portfolioMap[trade.Symbol].Amount = portfolioMap[trade.Symbol].Amount.Sub(trade.Amount)
				portfolioMap[trade.Symbol].AveragePrice = portfolioMap[trade.Symbol].AveragePrice.Sub(trade.FullPrice)
			}
		}

		var portfolio []Asset

		for _, asset := range portfolioMap {
			asset.AveragePrice = asset.AveragePrice.DivRound(asset.Amount, 3)
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
