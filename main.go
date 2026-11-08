package main

import (
	"net/http"
	"pokedex/internal/commands"
	"time"
)

func main() {
	client := http.Client{Timeout: 10 * time.Minute}
		
	commands.StartRepl(&client)
}
