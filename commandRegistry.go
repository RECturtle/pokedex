package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"map": {
			name:        "map",
			description: "Retrieve locations from the map",
			callback:    commandMap,
		},
		"bmap": {
			name:        "bmap",
			description: "Retrieve previous locations from the map",
			callback:    commandBmap,
		},
		"explore": {
			name:        "explore",
			description: "Retrieve pokemon in a location",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon's stats",
			callback:    inspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandCatch(config *Config) error {
	if len(config.Args) == 0 {
		return fmt.Errorf("Enter a pokemon to try and catch it")
	}
	name := config.Args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := config.Client.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("%s was caught!\n", name)
	config.Pokedex[name] = pokemon
	return nil
}

func commandExplore(config *Config) error {
	if len(config.Args) == 0 {
		return fmt.Errorf("Supply a city name from map to retrieve pokemon")
	}
	city := config.Args[0]

	pokemon, err := config.Client.GetLocationAreaPokemon(city)
	if err != nil {
		return err
	}
	for _, entry := range pokemon.PokemonEncounters {
		fmt.Println(entry.Pokemon.Name)
	}
	return nil
}

func commandMap(config *Config) error {
	if config.Next == "" {
		return fmt.Errorf("Out of entries!")
	}
	locations, err := config.Client.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, entry := range locations.Results {
		fmt.Println(entry.Name)
	}
	return nil
}

func commandBmap(config *Config) error {
	if config.Previous == "" {
		return fmt.Errorf("Out of entries!")
	}
	locations, err := config.Client.GetLocationAreas(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, entry := range locations.Results {
		fmt.Println(entry.Name)
	}
	return nil
}

func commandPokedex(config *Config) error {
	fmt.Println("Your Pokedex:")
	for name := range config.Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}

func inspectPokemon(config *Config) error {
	if len(config.Args) == 0 {
		return fmt.Errorf("Enter a pokemon you've caught to see stats")
	}
	name := config.Args[0]
	poke, ok := config.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name:", poke.Name)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weight:", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf(". -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf(". - %s\n", t.Type.Name)
	}
	return nil
}
