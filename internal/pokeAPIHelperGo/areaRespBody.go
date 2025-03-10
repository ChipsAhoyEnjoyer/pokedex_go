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
