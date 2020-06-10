package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func extractZipInMemory(data []byte) ([]byte, error){
	readerAt := bytes.NewReader(data)
	r, err := zip.NewReader(readerAt, int64(len(data)))
	if err != nil {
		return nil, err
	}
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		extractedData, err := ioutil.ReadAll(rc)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		rc.Close()
		return extractedData, nil
	}
	return nil, errors.New("no file found inside zip")
}

func downloadZip(year uint) {
	// Gets the price of a given ticker through B3's public API
	// Prices are 15 minutes in the past

	//url := fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A%v.ZIP", year)
	url := fmt.Sprintf("http://localhost:8000/COTAHIST_A2019.ZIP")
	response, err := http.Get(url)

	fmt.Println(response.Status)
	fmt.Println("--------")
	fmt.Println(response.Header)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	encoded_content, err := extractZipInMemory(responseData)
	fmt.Println(encoded_content)
	//header, content := ParseHistoricDataFromBytes(encoded_content)
	//fmt.Println(header)
	//fmt.Println("------")
	//fmt.Println(content)
	//fmt.Println("------")
}

func main() {
	fmt.Println("Starting...")
	downloadZip(2020)
	fmt.Println("Finishing")
}
