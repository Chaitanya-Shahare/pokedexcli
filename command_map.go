package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var Config = map[string]string{
	"Next":     "https://pokeapi.co/api/v2/location-area/",
	"Previous": "",
}

func callbackMap() error {

	res, err := http.Get(Config["Next"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	response := LocationsResponse{}

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
	}

	if response.Next != nil {
		Config["Next"] = *response.Next
	}
	if response.Previous != nil {
		Config["Previous"] = *response.Previous
	}

	fmt.Println()
	for _, result := range response.Results {
		fmt.Println("\t" + result.Name)
	}

	fmt.Println()

	return nil
}

func callbackMapb() error {
	if Config["Previous"] == "" {
		fmt.Println(Config["Previous"])
		fmt.Println("\tNo previous location areas")
		fmt.Println()
		return nil
	}

	res, err := http.Get(Config["Previous"])

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	response := LocationsResponse{}

	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
	}

	if response.Next != nil {

		Config["Next"] = *response.Next
	}
	if response.Previous != nil {
		Config["Previous"] = *response.Previous
	}

	fmt.Println()
	for _, result := range response.Results {
		fmt.Println("\t" + result.Name)
	}

	fmt.Println()

	return nil
}
