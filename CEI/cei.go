package main

import (
	"./base"
	"encoding/json"
	"fmt"
)

func getUserTrades(cpf, password string) string {
	trades := base.GetRawUserTrades(cpf, password)
	if trades != nil {
		tradesJsonString, err := json.Marshal(trades)
		if err == nil{
			return string(tradesJsonString)
		}else {
			return ""
		}
	} else {
		return ""
	}
}

func main() {
	fmt.Println(getUserTrades("14354069741", "@loucamente42"))
}