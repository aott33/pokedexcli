package pokeapi


type AreaResult struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type AreaResponse struct {
    Count    int          `json:"count"`
    Next     *string      `json:"next"`
    Previous *string      `json:"previous"`
    Results  []AreaResult `json:"results"`
}

type AreaPokemon struct {
    PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
    Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
    Name string `json:"name"`
}
