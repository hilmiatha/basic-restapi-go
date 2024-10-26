package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//struct dengan json tag
type FactResponse struct {
	Teks string `json:"text"`
	Tipe string `json:"type"`
}


func main() {
	// Create a new request
	req, err := http.NewRequest("GET", "https://cat-fact.herokuapp.com/facts/random", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// create a new client
	client := http.Client{}

	// send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// close the response body
	defer res.Body.Close()

	// read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// convert to type FactResponse
	var fact FactResponse = FactResponse{}
	json.Unmarshal(resBody, &fact)



	fmt.Printf("Type : %s - Facts : %s\n", fact.Tipe, fact.Teks)


	var factToSend FactResponse = FactResponse{"Kucing suka makan", "Kucing"}

	byteToSend, _ := json.Marshal(factToSend)

	fmt.Println(string(byteToSend))
	
}