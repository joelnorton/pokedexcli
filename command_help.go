package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, _ string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

	names := make([]string, 0, len(cfg.cliCommands))
	for name := range cfg.cliCommands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		c := cfg.cliCommands[name]
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}
