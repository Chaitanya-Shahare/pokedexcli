package main

import "fmt"

func callbackHelp() error {

	commands := GetCommands()

	fmt.Println("\tWelcome to the Pokedex!")
	fmt.Println()
	fmt.Println("\tUsage:")
	fmt.Println()

	for _, c := range commands {
		fmt.Printf("\t- %s: %s\n", c.name, c.description)
	}

	return nil
}
