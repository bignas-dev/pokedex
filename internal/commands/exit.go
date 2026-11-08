package commands

import (
	"os"
	"fmt"
	"net/http"
)

func commandExit(client *http.Client, input string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
