package main

import (
	"time"

	pokecache "github.com/joelnorton/pokedexcli/cache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	state       string
	cliCommands map[string]cliCommand
	next        string
	prev        string
	cache       *pokecache.Cache
	pokemon     map[string]Pokemon
}

func main() {
	cache := pokecache.NewCache(time.Minute * 5)
	startRepl(&config{
		cliCommands: map[string]cliCommand{
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
			"explore": {
				name:        "explore",
				description: "Explore an area in the world",
				callback:    getArea,
			},
			"catch": {
				name:        "catch",
				description: "Throw a pokeball at the pokemon",
				callback:    catchPokemon,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a Pokemon you have captured",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "View your Pokedex",
				callback:    commandPokedex,
			},
		},
		next:    "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		prev:    "",
		cache:   cache,
		pokemon: map[string]Pokemon{},
	})
}
