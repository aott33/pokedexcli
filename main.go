package main

import (
	"time"

	"github.com/aott33/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	
	cfg := &Config{
		PokeClient: client,
	}

	startRepl(cfg)

}
