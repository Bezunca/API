package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Header struct{
	TipReg int
	FileName string
	Source string
	GenDate time.Time
}

func check(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}

func parse_header(rawHeader string) Header{
	tipreg, err := strconv.ParseInt(rawHeader[:2], 10, 64)
	check(err)
	date, err := time.Parse("20060102", rawHeader[23:23+8])
	check(err)
	return Header{
		 int(tipreg),
		rawHeader[2:2+13],
		rawHeader[15:15+8],
		date,
	}
}

func main() {
	rawData, err := ioutil.ReadFile("../COTAHIST_A2019.TXT")
	check(err)
	segmentedLines := strings.Split(string(rawData), "\n")
	fmt.Println(segmentedLines[:2])
	fmt.Println("----------")
	header := parse_header(segmentedLines[0])
	fmt.Println(header)
}
