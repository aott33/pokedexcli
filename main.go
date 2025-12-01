package main

import (
	"math/rand"
	"time"

	"github.com/aott33/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second, 5 * time.Second)
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)


	cfg := &Config{
		PokeClient: client,
		Pokedex: 	make(map[string]pokeapi.PokemonInfo),
		RNG:		rng,
	}

	startRepl(cfg)

}
