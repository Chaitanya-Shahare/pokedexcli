package main

import "fmt"

func callbackInspect(args []string) error {

	fmt.Println()
	defer fmt.Println()

	if len(args) < 2 {
		fmt.Println("\tPlease provide the name of the Pokemon you caught.")
		return nil
	}

	pokemonName := args[1]

	pokemon, ok := Pokedex[pokemonName]

	if !ok {
		fmt.Println("\tYou have not caught a Pokemon named ", pokemonName)
		return nil
	}

	// Name: pidgey
	// Height: 3
	// Weight: 18
	// Stats:
	// -hp: 40
	// -attack: 45
	// -defense: 40
	// -special-attack: 35
	// -special-defense: 35
	// -speed: 56
	// Types:
	// - normal
	// - flying

	fmt.Println("\tName: ", pokemon.Name)
	fmt.Println("\tHeight: ", pokemon.Height)
	fmt.Println("\tWeight: ", pokemon.Weight)
	fmt.Println("\tStats:")
	for _, value := range pokemon.Stats {
		fmt.Printf("\t\t- %v: %v\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("\tTypes:")
	for _, t := range pokemon.Types {
		fmt.Printf("\t\t- %v\n", t.Type.Name)
	}

	return nil
}
