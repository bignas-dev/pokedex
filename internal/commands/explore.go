package commands

import (
	"fmt"
	"net/http"
)

type ExploreResponse struct {
	PokemonEncounters []PokemonEntry `json:"pokemon_encounters"`
}

type PokemonEntry struct {
	Entry Pokemon `json:"pokemon"`
}


type Pokemon struct {
	Name string `json:"name"`
}

func commandExplore(client *http.Client, input string) error {
	fmt.Printf("Exploring %v...\n", input)
	fmt.Println("Found Pokemon:")

	url := pokemonMapAPI + input

	var output ExploreResponse
	fetchData(client, url, &output)	

	for _, r := range output.PokemonEncounters {
		fmt.Printf("- %v\n", r.Entry.Name)
	}
	return nil

}
