package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func start_repl() {

	commands := GetCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		words := strings.Fields(text)

		command := words[0]

		c, ok := commands[command]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		c.callback()

	}
}
