package pokeapi

type PokemonInfo struct {
    BaseExperience int    `json:"base_experience"`
    Height         int    `json:"height"`
    Name           string `json:"name"`
    Stats          []Stats `json:"stats"`
    Types          []Types `json:"types"`
    Weight         int    `json:"weight"`
}

type Stat struct {
	Name string `json:"name"`
}
type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}
type Type struct {
	Name string `json:"name"`
}
type Types struct {
	Type Type `json:"type"`
}
