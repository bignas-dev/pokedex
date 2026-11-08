package commands

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

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


func commandCatch(client *http.Client, input string) error {
	url := pokemonPropertiesAPI + input
	fmt.Printf("Throwing a Pokeball at %v...\n", input)

	var output PokemonSpecies
	fetchData(client, url, &output)

	neededMin := float64(output.CaptureRate) / 255.0
	gotChance := rand.Float64()

	if gotChance <= neededMin {
		caughtPokemons[input] = output
		fmt.Printf("%v was caught!\n", input)
	} else {
		fmt.Printf("%v escaped!\n", input)
	}

	return nil
}

