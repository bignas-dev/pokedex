package commands

import (
	"fmt"
	"net/http"
)

func commandPokedex(client *http.Client, input string) error {
	fmt.Println("Your pokedex:")
	for pokemonName := range caughtPokemons {
		fmt.Printf("- %v\n", pokemonName)
	}
	return nil
}
