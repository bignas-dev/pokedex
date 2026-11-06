package main

import (
	"fmt"
	"time"
	"io"
	"encoding/json"
	"net/http"
	"pokedex/internal/pokecache"
)

var Cache = pokecache.NewCache(5 * time.Minute)

type MapResponse struct {
	Results []Names `json:"results"`
}

type Names struct {
	Name string `json:"name"`
}

const (
	apiUrl = "https://pokeapi.co/api/v2/location-area/"
	limit = 20
)

var offset = 0


func commandMap(cfg *config) error {
	url := fmt.Sprintf("%v?offset=%v&limit=%v", apiUrl, offset, limit)

	bytes, ok := Cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}

		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error getting response: %w", err)
		}
		defer resp.Body.Close()

		bytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error getting response: %w", err)
		}
		Cache.Add(url, bytes)
	} 

	var mapResp MapResponse
	if err := json.Unmarshal(bytes, &mapResp); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	for _, r := range mapResp.Results {
		fmt.Println(r.Name)
	}
	offset += limit
	return nil
}

func commandMapb(cfg *config) error {
	offset = max(0, offset - limit - limit)
	url := fmt.Sprintf("%v?offset=%v&limit=%v", apiUrl, offset, limit)

	bytes, ok := Cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}

		client := &http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error getting response: %w", err)
		}
		defer resp.Body.Close()

		bytes, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error getting response: %w", err)
		}
		Cache.Add(url, bytes)
	} 

	var mapResp MapResponse
	if err := json.Unmarshal(bytes, &mapResp); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	for _, r := range mapResp.Results {
		fmt.Println(r.Name)
	}
	offset += limit
	return nil
}
