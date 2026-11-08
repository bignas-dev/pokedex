package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type config struct{}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, secondval string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Gets next 20 map elements",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets previous 20 map elements",
			callback:    commandMapb,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows list of caught pokemons",
			callback:    commandPokedex,
		},
		"explore": {
			name:        "explore",
			description: "Explores a given location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			log.Fatalf("error scanner.Scan(): %v", scanner.Err())
		}

		input := scanner.Text()
		clnInpt := cleanInput(input)
		cmdName := clnInpt[0]
		argName := ""
		if len(clnInpt) > 1 {
			argName = clnInpt[1]
		}
		if cmdInfo, ok := getCommands()[cmdName]; ok {
			if err := cmdInfo.callback(cfg, argName); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
