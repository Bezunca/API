package main

import (
	"./scrapper"
	"encoding/json"
	"fmt"
)

func main() {

	/*trades := scrapper.GetUserTrades("CPF", "SENHA")
	tradesJson, _ := json.Marshal(trades)
	fmt.Println("\n\n\n", string(tradesJson))*/

	/*portfolio := scrapper.GetUserPortfolioFromTrades("CPF", "SENHA")
	portfolioJson, _ := json.Marshal(portfolio)
	fmt.Println("\n\n\n", string(portfolioJson))*/

	/*dividends := scrapper.GetUserDividends("CPF", "SENHA")
	dividendsJson, _ := json.Marshal(dividends)
	fmt.Println("\n\n\n", string(dividendsJson))*/

	portfolio := scrapper.GetUserPortfolio("CPF", "SENHA")
	portfolioJson, _ := json.Marshal(portfolio)
	fmt.Println("\n\n\n", string(portfolioJson))
}
