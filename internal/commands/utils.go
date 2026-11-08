// Package commands contains the commands and the repl
package commands

import (
	"time"
	"io"
	"fmt"
	"encoding/json"
	"net/http"
	"pokedex/internal/pokecache"
)

const (
	pokemonPropertiesAPI = "https://pokeapi.co/api/v2/pokemon-species/"
	pokemonMapAPI = "https://pokeapi.co/api/v2/location-area/"
	pokemonMapAPILimit = 20
)

var caughtPokemons = map[string]PokemonSpecies{}

var cache = pokecache.NewCache(5 * time.Minute)
var offset = 0

func fetchData[T any](client *http.Client, url string, output *T) error {
	bytes, ok := cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
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
		cache.Add(url, bytes)
	} 

	if err := json.Unmarshal(bytes, &output); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}

