package main

import (
	"fmt"
	"os"

	"github.com/aott33/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	PokeClient	pokeapi.Client
	Next		*string
	Previous	*string
}


func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
		"exit": {
        	name:        "exit",
        	description: "Exit the Pokedex",
        	callback:    commandExit,
    	},
		"help": {
			name:		 "help",
			description: "Displays a help message",
			callback:	 commandHelp,
		},
		"map": {
			name:		 "map",
			description: "Displays the next 20 location areas",
			callback: 	 commandMap,
		},
		"mapb": {
			name: 		 "mapb",
			description: "Display the previous 20 location areas",
			callback: 	 commandMapb,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(cfg *Config) error {
	url := pokeapi.BaseURL + "/location-area"

	if cfg.Next != nil {
		url = *cfg.Next
	}

	areaResp, err := cfg.PokeClient.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, r := range areaResp.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = areaResp.Next
	cfg.Previous = areaResp.Previous

	return nil
}


func commandMapb(cfg *Config) error {

	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.Previous

	areaResp, err := cfg.PokeClient.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, r := range areaResp.Results {
		fmt.Println(r.Name)
	}

	cfg.Next = areaResp.Next
	cfg.Previous = areaResp.Previous

	return nil
}
