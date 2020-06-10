package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	TipReg   int
	FileName string
	Source   string
	GenDate  time.Time
}

type FixedPoint int // Represents the actual value multiplied by 100

type SecurityQuote struct {
	TipReg                 int
	DataCollectionDate     time.Time
	BDICode                int
	Ticker                 string
	MarketType             int
	CompanyName            string
	SecurityType           string
	FutureMarketExpiration string // Need to check this against an example that actually has a value
	Currency               string
	PriceOpen              FixedPoint
	PriceMax               FixedPoint
	PriceMin               FixedPoint
	PriceMean              FixedPoint
	PriceClose             FixedPoint
	PriceBid               FixedPoint
	PriceAsk               FixedPoint
	TotalTrades            int
	TotalQuantity          int
	TotalVolume            FixedPoint
	PreExe                 FixedPoint // Needs further investigation
	IndOpc                 int        // Needs further investigation
	ExpirationDate         time.Time
	FatCot                 int // Needs further investigation
	PtoExe                 int // Needs further investigation
	ISINCode               string
	DistributionNumber     int
}

func check(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func parseHeader(rawHeader string) Header {
	tipreg, err := strconv.ParseInt(rawHeader[:2], 10, 64)
	check(err)
	date, err := time.Parse("20060102", rawHeader[23:23+8])
	check(err)
	return Header{
		int(tipreg),
		rawHeader[2 : 2+13],
		rawHeader[15 : 15+8],
		date,
	}
}

func parseContentLine(raw_line string) SecurityQuote {
	tipReg, err := strconv.ParseInt(raw_line[:2], 10, 64)
	check(err)
	date, err := time.Parse("20060102", raw_line[2:2+8])
	check(err)
	bdiCode, err := strconv.ParseInt(raw_line[10:10+2], 10, 64)
	check(err)
	marketType, err := strconv.ParseInt(raw_line[24:24+3], 10, 64)
	check(err)
	priceOpen, err := strconv.ParseInt(raw_line[56:56+13], 10, 64)
	check(err)
	priceMax, err := strconv.ParseInt(raw_line[69:69+13], 10, 64)
	check(err)
	priceMin, err := strconv.ParseInt(raw_line[82:82+13], 10, 64)
	check(err)
	priceMean, err := strconv.ParseInt(raw_line[95:95+13], 10, 64)
	check(err)
	priceClose, err := strconv.ParseInt(raw_line[108:108+13], 10, 64)
	check(err)
	priceBid, err := strconv.ParseInt(raw_line[121:121+13], 10, 64)
	check(err)
	priceAsk, err := strconv.ParseInt(raw_line[134:134+13], 10, 64)
	check(err)
	totalTrades, err := strconv.ParseInt(raw_line[147:147+5], 10, 64)
	check(err)
	totalQuantity, err := strconv.ParseInt(raw_line[152:152+18], 10, 64)
	check(err)
	totalVolume, err := strconv.ParseInt(raw_line[170:170+18], 10, 64)
	check(err)
	preExe, err := strconv.ParseInt(raw_line[188:188+13], 10, 64)
	check(err)
	indOpc, err := strconv.ParseInt(raw_line[201:201+1], 10, 64)
	check(err)
	expirationDate, err := time.Parse("20060102", raw_line[202:202+8])
	check(err)
	fatCot, err := strconv.ParseInt(raw_line[210:210+7], 10, 64)
	check(err)
	ptoExe, err := strconv.ParseInt(raw_line[210:210+7], 10, 64)
	check(err)
	distributionNumber, err := strconv.ParseInt(raw_line[242:242+3], 10, 64)
	check(err)
	return SecurityQuote{
		int(tipReg),
		date,
		int(bdiCode),
		strings.TrimSpace(raw_line[12 : 12+12]),
		int(marketType),
		strings.TrimSpace(raw_line[27 : 27+12]),
		strings.TrimSpace(raw_line[39 : 39+10]),
		strings.TrimSpace(raw_line[49 : 49+3]),
		strings.TrimSpace(raw_line[52 : 52+4]),
		FixedPoint(priceOpen),
		FixedPoint(priceMax),
		FixedPoint(priceMin),
		FixedPoint(priceMean),
		FixedPoint(priceClose),
		FixedPoint(priceBid),
		FixedPoint(priceAsk),
		int(totalTrades),
		int(totalQuantity),
		FixedPoint(totalVolume),
		FixedPoint(preExe),
		int(indOpc),
		expirationDate,
		int(fatCot),
		int(ptoExe),
		strings.TrimSpace(raw_line[230 : 230+12]),
		int(distributionNumber),
	}
}

func ParseHistoricData(rawData []string) (Header, []SecurityQuote) {
	header := parseHeader(rawData[0])
	contentList := make([]SecurityQuote, len(rawData)-3)
	for i, raw_line := range rawData[1:len(rawData)-2] {
		contentList[i] = parseContentLine(raw_line)
	}
	return header, contentList
}

func ParseHistoricDataFromBytes(data []byte) (Header, []SecurityQuote){
	segmentedLines := strings.Split(string(data), "\n")
	return ParseHistoricData(segmentedLines)
}

func main() {
	rawData, err := ioutil.ReadFile("../COTAHIST_A2019.TXT")
	check(err)
	header, contents := ParseHistoricDataFromBytes(rawData)
	fmt.Println(header)
	fmt.Println("------")
	fmt.Println(contents)
	fmt.Println("------")
}
