package main

import (
	"strings"
	"os"
	"fmt"
	"log"
	"bufio"
)

type config struct {}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
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

		if cmdInfo, ok := getCommands()[cmdName]; ok {
			if err := cmdInfo.callback(cfg); err != nil {
				log.Fatal(err)
			}

		} else {
			fmt.Println("Unknown command")
		}
	}
}
