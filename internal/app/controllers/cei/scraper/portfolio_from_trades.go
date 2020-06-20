package scraper

import (
	"bezuncapi/internal/models"
	"github.com/shopspring/decimal"
)

func GetUserPortfolioFromTrades(cpf, password string) []models.Asset {
	trades := GetUserTrades(cpf, password)

	if len(trades) > 0 {

		portfolioMap := map[string]*models.Asset{}

		for _, trade := range trades {

			if _, ok := portfolioMap[trade.Symbol]; !ok {
				portfolioMap[trade.Symbol] = &models.Asset{
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

		var portfolio []models.Asset

		for _, asset := range portfolioMap {
			asset.AveragePrice = asset.AveragePrice.DivRound(asset.Amount, 3)
			portfolio = append(portfolio, *asset)
		}
		return portfolio
	} else {
		return nil
	}
}
