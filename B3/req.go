package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func main() {
	var data map[string]interface{}
	response, err := http.Get("https://arquivos.b3.com.br/apinegocios/ticker/ITSA4/2020-05-12")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println("------")
	values, ok := data["innermap"].([]interface{})
	fmt.Println(values[0])
}