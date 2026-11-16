package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func commandNext(cfg *config) error {
	res, err := http.Get(cfg.next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	locations := Location{}
	err = json.Unmarshal(body, &locations)
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

func commandPrev(cfg *config) error {
	if cfg.prev == "" {
		return fmt.Errorf("No previous map found")
	}
	res, err := http.Get(cfg.prev)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	locations := Location{}
	err = json.Unmarshal(body, &locations)
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
