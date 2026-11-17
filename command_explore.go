package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type AreaElements struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, location string) error {
	u, err := url.JoinPath("https://pokeapi.co/api/v2/location-area/", url.PathEscape(location))
	if err != nil {
		return err
	}

	v, exists := cfg.cache.GetCache(u)

	area := AreaElements{}

	if exists {
		err = json.Unmarshal(v, &area)
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

		err = json.Unmarshal(body, &area)
		if err != nil {
			return err
		}
	}
	fmt.Println("Pokemon:")
	for _, e := range area.PokemonEncounters {
		fmt.Println(" - " + e.Pokemon.Name)
	}

	return nil
}
