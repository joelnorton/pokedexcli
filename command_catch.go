package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
)

func catchPokemon(cfg *config, pokemon string) error {
	u, err := url.JoinPath("https://pokeapi.co/api/v2/pokemon/", url.PathEscape(pokemon))
	if err != nil {
		return err
	}

	key := pokemon
	v, exists := cfg.cache.GetCache(key)

	Pokemon := Pokemon{}

	if exists {
		err = json.Unmarshal(v, &Pokemon)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(u)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}

		cfg.cache.AddCache(key, body)

		err = json.Unmarshal(body, &Pokemon)
		if err != nil {
			return err
		}
	}
	if Pokemon.BaseExperience < 1 {
		return fmt.Errorf("Pokemon experience error")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	number := rand.IntN(Pokemon.BaseExperience)

	if (Pokemon.BaseExperience/3) < number && number > (Pokemon.BaseExperience/3*2) {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokemon[pokemon] = Pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
