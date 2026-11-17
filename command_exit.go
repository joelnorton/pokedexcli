package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	cfg.cache.Close()
	os.Exit(0)
	return nil
}
