package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/recturtle/pokedexcli/internal/pokeapi"
)

type Config struct {
	Client   *pokeapi.Client
	Pokedex  map[string]pokeapi.Pokemon
	Args     []string
	Next     string
	Previous string
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	trimmed := strings.TrimSpace(lowered)
	split := strings.Fields(trimmed)
	return split
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(pokeapi.DefaultBaseURL)
	config := &Config{
		Client:  client,
		Pokedex: make(map[string]pokeapi.Pokemon),
		Next:    client.InitialLocationAreaURL(),
	}
	commands := getCommands()

	for {
		var input string
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input = scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		res, ok := commands[cleaned[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		config.Args = cleaned[1:]

		if err := res.callback(config); err != nil {
			fmt.Println(err)
		}
	}
}
