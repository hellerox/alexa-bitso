package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getBitsoPrice(book string) payload {
	url := fmt.Sprintf("https://api.bitso.com/v3/ticker/")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)

	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)

	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record priceRequest
	// Use json.Decode for reading streams of JSON data
	if errg := json.NewDecoder(resp.Body).Decode(&record); errg != nil {
		log.Println(errg)
	}

	for _, v := range record.Payload {
		log.Printf("%s y %s", book, v.Book)
		if book == v.Book {
			return v
		}
	}

	log.Printf("%+v", record)
	return record.Payload[0]
}

func getBitsoResponse(book string) (bitsoResponse string) {
	price := getBitsoPrice(book)
	log.Printf("resultado para book: %s es %+v", book, price)
	bitsoResponse = fmt.Sprintf("el precio actual de %s es %s", pairDescription[price.Book], price.Last)
	return
}
