package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Location struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func commandMap(cfg *config, _ string) error {
	locations, err := getMap(cfg, cfg.next)
	if err != nil {
		return err
	}
	cfg.next = locations.Next
	cfg.prev = locations.Previous

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapb(cfg *config, _ string) error {
	if cfg.prev == "" {
		return fmt.Errorf("No previous map found\n")
	}
	locations, err := getMap(cfg, cfg.prev)
	if err != nil {
		return err
	}
	cfg.next = locations.Next
	cfg.prev = locations.Previous

	for _, l := range locations.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func getMap(cfg *config, u string) (Location, error) {
	locations := Location{}
	parsed, err := url.Parse(u)
	if err != nil {
		return locations, err
	}

	key := parsed.String()
	v, exists := cfg.cache.GetCache(key)
	if exists {
		err = json.Unmarshal(v, &locations)
		if err != nil {
			return locations, err
		}
	} else {
		res, err := httpClient.Get(u)
		if err != nil {
			return locations, err
		}
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return locations, fmt.Errorf("unexpected status from PokeAPI: %s", res.Status)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return locations, err
		}

		cfg.cache.AddCache(key, body)

		err = json.Unmarshal(body, &locations)
		if err != nil {
			return locations, err
		}
	}
	return locations, nil
}
