package main

import (
	"encoding/json"
	"fmt"
	"github.com/Chaitanya-Shahare/pokedexcli/pokecache"
	"io"
	"net/http"
	"time"
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

var cache = pokecache.NewCache(5 * time.Second)

func callbackMap(args []string) error {
	data, err := fetchData(Config["Next"])
	if err != nil {
		return err
	}

	if data, ok := cache.Get(Config["Next"]); ok {
		return processLocationsResponse(data, Config)
	}

	cache.Add(Config["Next"], data)

	return processLocationsResponse(data, Config)
}

func callbackMapb(args []string) error {
	if Config["Previous"] == "" {
		fmt.Println(Config["Previous"])
		fmt.Println("\tNo previous location areas")
		fmt.Println()
		return nil
	}

	if data, ok := cache.Get(Config["Previous"]); ok {
		return processLocationsResponse(data, Config)
	}

	data, err := fetchData(Config["Previous"])
	if err != nil {
		return err
	}

	cache.Add(Config["Previous"], data)

	return processLocationsResponse(data, Config)
}

func fetchData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	return body, nil
}

func processLocationsResponse(data []byte, config map[string]string) error {
	response := LocationsResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return err
	}

	if response.Next != nil {
		config["Next"] = *response.Next
	}
	if response.Previous != nil {
		config["Previous"] = *response.Previous
	}

	fmt.Println()
	for _, result := range response.Results {
		fmt.Println("\t" + result.Name)
	}
	fmt.Println()

	return nil
}
