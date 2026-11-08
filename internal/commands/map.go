package commands

import (
	"fmt"
	"net/http"
)


type MapResponse struct {
	Results []Names `json:"results"`
}

type Names struct {
	Name string `json:"name"`
}


func commandMap(client *http.Client, input string) error {
	url := fmt.Sprintf("%v?offset=%v&limit=%v", pokemonMapAPI, offset, pokemonMapAPILimit)

	var output MapResponse
	fetchData(client, url, &output)

	for _, r := range output.Results {
		fmt.Println(r.Name)
	}

	offset += pokemonMapAPILimit
	return nil
}

func commandMapb(client *http.Client, input string) error {
	offset = max(0, offset - pokemonMapAPILimit - pokemonMapAPILimit)
	url := fmt.Sprintf("%v?offset=%v&limit=%v", pokemonMapAPI, offset, pokemonMapAPILimit)

	var output MapResponse
	fetchData(client , url , &output )

	for _, r := range output.Results {
		fmt.Println(r.Name)
	}
	offset += pokemonMapAPILimit
	return nil
}

