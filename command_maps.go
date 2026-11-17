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
		Url  string `json:"url"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func commandNext(cfg *config, _ string) error {
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

func commandPrev(cfg *config, _ string) error {
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
	url, err := url.Parse(u)
	if err != nil {
		return locations, err
	}

	key := url.Query().Get("offset")
	v, exists := cfg.cache.GetCache(key)
	if exists {
		err = json.Unmarshal(v, &locations)
		if err != nil {
			return locations, err
		}
	} else {
		res, err := http.Get(u)
		if err != nil {
			return locations, err
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
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
