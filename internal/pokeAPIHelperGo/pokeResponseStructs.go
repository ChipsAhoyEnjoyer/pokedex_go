package pokeAPIHelperGo

import "github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"

type LocationAreas struct {
	Cache  *pokeCache.PokeCache
	Next   string `json:"next"`
	Prev   string `json:"previous"`
	Result []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type PokeEncounters struct {
	Cache        *pokeCache.PokeCache
	LocationName string `json:"name"`
	Encounters   []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	LocationAreaEncounters string `json:"location_area_encounters"`
}
