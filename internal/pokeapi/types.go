package pokeapi

// https://pokeapi.co/api/v2/pokemon/
type Pokemon struct {
	Name           string        `json:"name"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	BaseExperience int           `json:"base_experience"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonType struct {
	Type PokeTypeInfo `json:"type"`
}

type PokeTypeInfo struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	BaseStat int             `json:"base_stat"`
	Effort   int             `json:"effort"`
	Stat     PokemonStatInfo `json:"stat"`
}

type PokemonStatInfo struct {
	Name string `json:"name"`
}

type LocationAreaEntry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []LocationAreaEntry `json:"results"`
}

type LocationAreaDetailResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonEntry `json:"pokemon"`
}

type PokemonEntry struct {
	Name string `json:"name"`
}
