package main

import "fmt"

func commandHelp(cfg *config, _ string) error {

	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, c := range cfg.cliCommands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
