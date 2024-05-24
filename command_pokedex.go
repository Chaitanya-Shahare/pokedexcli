package main

import "fmt"

func callbackPokedex(args []string) error {
	fmt.Println()
	defer fmt.Println()

	fmt.Println("\tYour Pokedex:")

	for name, _ := range Pokedex {
		fmt.Printf("\t\t- %v\n", name)
	}

	return nil
}
