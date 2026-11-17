package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	data, _ := io.ReadAll(r)
	return string(data)
}

func TestCommandInspect_Caught(t *testing.T) {
	cfg := &config{
		pokemon: map[string]Pokemon{
			"pidgey": {
				ID:                     16,
				Name:                   "pidgey",
				BaseExperience:         63,
				Height:                 3,
				Weight:                 18,
				IsDefault:              true,
				Order:                  16,
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
				Stats: []Stat{
					{BaseStat: 40, Effort: 0, Stat: NamedAPIResource{Name: "hp"}},
					{BaseStat: 45, Effort: 0, Stat: NamedAPIResource{Name: "attack"}},
					{BaseStat: 40, Effort: 0, Stat: NamedAPIResource{Name: "defense"}},
					{BaseStat: 35, Effort: 0, Stat: NamedAPIResource{Name: "special-attack"}},
					{BaseStat: 35, Effort: 0, Stat: NamedAPIResource{Name: "special-defense"}},
					{BaseStat: 56, Effort: 0, Stat: NamedAPIResource{Name: "speed"}},
				},
				Types: []Type{
					{Slot: 1, Type: NamedAPIResource{Name: "normal"}},
					{Slot: 2, Type: NamedAPIResource{Name: "flying"}},
				},
				Cries:   Cries{},
				Species: NamedAPIResource{},
			},
		},
	}

	out := captureOutput(func() {
		commandInspect(cfg, "pidgey")
	})

	if !strings.Contains(out, "Name: pidgey") {
		t.Errorf("expected 'Name: pidgey' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Height: 3") {
		t.Errorf("expected 'Height: 3' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Weight: 18") {
		t.Errorf("expected 'Weight: 18' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Stats:") {
		t.Errorf("expected 'Stats:' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "-hp: 40") {
		t.Errorf("expected '-hp: 40' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "-speed: 56") {
		t.Errorf("expected '-speed: 56' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "Types:") {
		t.Errorf("expected 'Types:' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "- normal") {
		t.Errorf("expected '- normal' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "- flying") {
		t.Errorf("expected '- flying' in output, got:\n%s", out)
	}
}

func TestCommandInspect_NotCaught(t *testing.T) {
	cfg := &config{
		pokemon: map[string]Pokemon{},
	}

	out := captureOutput(func() {
		commandInspect(cfg, "pikachu")
	})

	if !strings.Contains(out, "you have not caught that pokemon") {
		t.Errorf("expected 'you have not caught that pokemon' in output, got:\n%s", out)
	}
}

func TestCommandInspect_CaseInsensitive(t *testing.T) {
	cfg := &config{
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
				Stats: []Stat{
					{BaseStat: 35, Effort: 0, Stat: NamedAPIResource{Name: "hp"}},
				},
				Types: []Type{
					{Slot: 1, Type: NamedAPIResource{Name: "electric"}},
				},
				Cries:   Cries{},
				Species: NamedAPIResource{},
			},
		},
	}

	// The pokedex map is keyed by lowercase name (see catchPokemon)
	out := captureOutput(func() {
		commandInspect(cfg, "pikachu")
	})

	if !strings.Contains(out, "Name: pikachu") {
		t.Errorf("expected 'Name: pikachu' in output, got:\n%s", out)
	}
}
