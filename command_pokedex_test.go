package main

import (
	"strings"
	"testing"
	"time"

	pokecache "github.com/joelnorton/pokedexcli/cache"
)

func TestCommandPokedex_Empty(t *testing.T) {
	cfg := &config{
		cliCommands: map[string]cliCommand{
			"help": {"help", "Displays help", commandHelp},
			"exit": {"exit", "Exit", commandExit},
		},
		cache:   pokecache.NewCache(time.Minute * 5),
		pokemon: map[string]Pokemon{},
	}

	out := captureOutput(func() {
		commandPokedex(cfg, "")
	})

	if !strings.Contains(out, "You have not captured any pokemon") {
		t.Errorf("expected empty-pokedex message, got %q", out)
	}
}

func TestCommandPokedex_WithPokemon(t *testing.T) {
	cfg := &config{
		cliCommands: map[string]cliCommand{
			"help": {"help", "Displays help", commandHelp},
			"exit": {"exit", "Exit", commandExit},
		},
		cache: pokecache.NewCache(time.Minute * 5),
		pokemon: map[string]Pokemon{
			"pikachu": {
				ID:                     25,
				Name:                   "pikachu",
				BaseExperience:         112,
				Height:                 4,
				Weight:                 60,
				IsDefault:              true,
				Order:                  25,
				Abilities:              []Ability{},
				PastAbilities:          []PastAbility{},
				Forms:                  []Form{},
				GameIndices:            []GameIndex{},
				HeldItems:              []HeldItem{},
				LocationAreaEncounters: "",
				Moves:                  []Move{},
				PastStats:              []PastStat{},
				PastTypes:              []PastType{},
				Sprites:                Sprites{},
				Stats:                  []Stat{},
				Types:                  []Type{},
				Cries:                  Cries{},
				Species:                NamedAPIResource{},
			},
		},
	}

	out := captureOutput(func() {
		commandPokedex(cfg, "")
	})

	if !strings.Contains(out, "Your Pokedex:") {
		t.Errorf("expected 'Your Pokedex:' line, got %q", out)
	}
	if !strings.Contains(out, "pikachu") {
		t.Errorf("expected 'pikachu' in list, got %q", out)
	}
}
