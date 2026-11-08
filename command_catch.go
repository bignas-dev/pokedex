package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
)

var caughtPokemons = map[string]PokemonSpecies{}

type PokemonSpecies struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Order                int    `json:"order"`
	GenderRate           int    `json:"gender_rate"`
	CaptureRate          int    `json:"capture_rate"`
	BaseHappiness        int    `json:"base_happiness"`
	IsBaby               bool   `json:"is_baby"`
	IsLegendary          bool   `json:"is_legendary"`
	IsMythical           bool   `json:"is_mythical"`
	HatchCounter         int    `json:"hatch_counter"`
	HasGenderDifferences bool   `json:"has_gender_differences"`
	FormsSwitchable      bool   `json:"forms_switchable"`
}

const (
	pokemonAPI = "https://pokeapi.co/api/v2/pokemon-species/"
)

func commandCatch(cfg *config, pokemonName string) error {
	url := pokemonAPI + pokemonName
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("http new request messed up: %w", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http response messed up: %w", err)
	}

	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error getting response: %w", err)
	}

	var mapResp PokemonSpecies

	if err := json.Unmarshal(bytes, &mapResp); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	neededMin := float64(mapResp.CaptureRate) / 255.0
	gotChance := rand.Float64()

	if gotChance <= neededMin {
		caughtPokemons[pokemonName] = mapResp
		fmt.Printf("%v was caught!\n", pokemonName)
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(cfg *config, pokemonName string) error {
	if val, ok := caughtPokemons[pokemonName]; ok {
		fmt.Println("- ID: ", val.ID)
		fmt.Println("- Name: ", val.Name)
		fmt.Println("- Order:", val.Order)
		fmt.Println("- Gender Rate:", val.GenderRate)
		fmt.Println("- Capture Rate: ", val.CaptureRate)
		fmt.Println("- Base, Happiness: ", val.BaseHappiness)
		fmt.Println("- Is Baby: ", val.IsBaby)
		fmt.Println("- Is IsLegendary: ", val.IsLegendary)
		fmt.Println("- Is Mythical: ", val.IsMythical)
		fmt.Println("- Hatch Counter: ", val.HatchCounter)
		fmt.Println(" - Has Gender Differences: ", val.HasGenderDifferences)
		fmt.Println("- Forms Switchable: ", val.FormsSwitchable)
	} else {
		fmt.Println("you have not caught this pokemon")
	}
	return nil
}

func commandPokedex(cfg *config, secondarg string) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range caughtPokemons {
		fmt.Printf("- %v\n", pokemon.Name)
	}
	return nil
}
