package main

import "fmt"

func commandPokedex(cfg *config, _ string) error {
	if len(cfg.pokemon) > 0 {
		fmt.Println("Your Pokedex:")
		for _, p := range cfg.pokemon {
			fmt.Printf(" - %s\n", p.Name)
		}
	} else {
		fmt.Println("You have not captured any pokemon")
	}
	return nil
}
