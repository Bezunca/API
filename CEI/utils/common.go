package utils

import (
	"github.com/shopspring/decimal"
	"log"
	"os"
	"regexp"
	"strings"
)

func CleanString(s string) string {
	if s != "" {
		space := regexp.MustCompile(`\s+`)
		return space.ReplaceAllString(strings.ReplaceAll(strings.TrimSpace(s), "\n", ""), " ")
	} else {
		return ""
	}
}

func StringToDecimal(s string) decimal.Decimal {
	s = strings.ReplaceAll(s, ",", ".")
	sInt, err := decimal.NewFromString(s)
	Check(err)

	return sInt
}

func Check(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
