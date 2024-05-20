package main

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the names of NEXT 20 location areas in the Pokemon world",
			callback:    callbackMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the names of PREVIOUS 20 location areas in the Pokemon world",
			callback:    callbackMapb,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}