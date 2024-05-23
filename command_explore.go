package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Chaitanya-Shahare/pokedexcli/pokecache"
)

type LocationResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var exploreCache = pokecache.NewCache(5 * time.Second)

func callbackExplore(args []string) error {
	if len(args) < 2 {
		fmt.Println()
		fmt.Println("\tPlease provide a location area ID")
		fmt.Println()
		return nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + args[1]

	if data, ok := exploreCache.Get(url); ok {
		var location LocationResponse
		err := json.Unmarshal(data, &location)
		if err != nil {
			return err
		}
		for _, encounter := range location.PokemonEncounters {
			fmt.Printf("\t  - %s\n", encounter.Pokemon.Name)
		}
		return nil
	}

	data, err := fetchData(url)
	if err != nil {
		return err
	}

	exploreCache.Add(url, data)

	var location LocationResponse
	err = json.Unmarshal(data, &location)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("\t  - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
