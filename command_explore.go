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
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func getArea(cfg *config, location string) error {
	u, err := url.JoinPath("https://pokeapi.co/api/v2/location-area/", url.PathEscape(location))
	if err != nil {
		return err
	}

	key := location
	v, exists := cfg.cache.GetCache(key)

	AreaElements := AreaElements{}

	if exists {
		err = json.Unmarshal(v, &AreaElements)
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

		err = json.Unmarshal(body, &AreaElements)
		if err != nil {
			return err
		}
	}
	fmt.Println("Pokemon:")
	for _, e := range AreaElements.PokemonEncounters {
		fmt.Println(" - " + e.Pokemon.Name)
	}

	return nil
}
