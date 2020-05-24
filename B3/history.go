package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func openZipFromBytes(data []byte) {
	dataReader := io.ReaderAt(data, 0)
}

func downloadZip(year uint) {
	// Gets the price of a given ticker through B3's public API
	// Prices are 15 minutes in the past

	//url := fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A%v.ZIP", year)
	url := fmt.Sprintf("http://localhost:8000/COTAHIST_A2020.ZIP")
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

	//defer response.Body.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	zipFromBytes := openZipFromBytes()

	for _, f := range response.Body.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}

	fmt.Println("---------")
	fmt.Println()
}

func main() {
	fmt.Println("Starting...")
	downloadZip(2020)
	fmt.Println("Finishing")
}
