package main

import (
	"./scrapper"
	"github.com/shopspring/decimal"
)

func getUserTrades(cpf, password string) []scrapper.Trade {
	return scrapper.GetUserTrades(cpf, password)
}

func getUserDividends(cpf, password string) map[string][]scrapper.Dividend {
	return scrapper.GetUserDividends(cpf, password)
}

func getUserPortfolio(cpf, password string) []scrapper.Asset {
	return scrapper.GetUserPortfolio(cpf, password)
}

func getUserPortfolioFromTrades(cpf, password string) []scrapper.Asset {
	trades := scrapper.GetUserTrades(cpf, password)

	if len(trades) > 0 {

		portfolioMap := map[string]*scrapper.Asset{}

		for _, trade := range trades {

			if _, ok := portfolioMap[trade.Symbol]; !ok {
				portfolioMap[trade.Symbol] = &scrapper.Asset{
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

		var portfolio []scrapper.Asset

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

	/*portfolio := getUserPortfolioFromTrades("CPF", "SENHA")
	portfolioJson, _ := json.Marshal(portfolio)
	fmt.Println("\n\n\n", string(portfolioJson))*/

	/*dividends := getUserDividends("CPF", "SENHA")
	dividendsJson, _ := json.Marshal(dividends)
	fmt.Println("\n\n\n", string(dividendsJson))*/

	/*portfolio := getUserPortfolio("CPF", "SENHA")
	portfolioJson, _ := json.Marshal(portfolio)
	fmt.Println("\n\n\n", string(portfolioJson))*/
}
