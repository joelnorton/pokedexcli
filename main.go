package main

func main() {
	startRepl(&config{cliCommands: map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 location ares in the Pokemon world",
			callback:    commandNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 location ares in the Pokemon world",
			callback:    commandPrev,
		},
	},
		next: "https://pokeapi.co/api/v2/location-area",
		prev: "",
	})
}
