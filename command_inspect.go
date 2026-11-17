package main

import "fmt"

func commandInspect(cfg *config, pokemon string) error {
	p, ok := cfg.pokemon[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	/*Name: pidgey
		Height: 3
		Weight: 18
		Stats:
	  -hp: 40
	  -attack: 45
	  -defense: 40
	  -special-attack: 35
	  -special-defense: 35
	  -speed: 56
		Types:
	  - normal
	  - flying */

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
