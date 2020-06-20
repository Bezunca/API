package models

import "github.com/shopspring/decimal"

type Trade struct {
	Date        string
	Action      string
	Market      string
	Expiration  string
	Symbol      string
	Description string
	Amount      decimal.Decimal
	Price       decimal.Decimal
	FullPrice   decimal.Decimal
	PriceFactor decimal.Decimal
}

type Asset struct {
	Symbol       string
	Description  string
	Market       string
	Amount       decimal.Decimal
	AveragePrice decimal.Decimal
}

type Dividend struct {
	Description  string
	Symbol       string
	Date         string
	Type         string
	BaseQuantity decimal.Decimal
	PriceFactor  decimal.Decimal
	GrossIncome  decimal.Decimal
	NetIncome    decimal.Decimal
}
