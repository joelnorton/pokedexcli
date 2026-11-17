package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"
)

func commandCatch(cfg *config, pokemon string) error {
	u, err := url.JoinPath("https://pokeapi.co/api/v2/pokemon/", url.PathEscape(pokemon))
	if err != nil {
		return err
	}

	v, exists := cfg.cache.GetCache(u)

	p := Pokemon{}

	if exists {
		err = json.Unmarshal(v, &p)
		if err != nil {
			return err
		}
	} else {
		res, err := httpClient.Get(u)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status from PokeAPI: %s", res.Status)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.AddCache(u, body)

		err = json.Unmarshal(body, &p)
		if err != nil {
			return err
		}
	}
	if p.BaseExperience < 1 {
		return fmt.Errorf("Pokemon experience error")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	number := rand.IntN(p.BaseExperience)

	if (p.BaseExperience/3) < number && number > (p.BaseExperience/3*2) {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokemon[pokemon] = p
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
