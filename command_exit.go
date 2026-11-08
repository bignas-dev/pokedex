package main

import (
	"os"
	"fmt"
)

func commandExit(cfg *config, secondval string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
