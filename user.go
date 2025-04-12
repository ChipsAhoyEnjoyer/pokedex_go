package main

import (
	"time"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"
)

const base_url = "https://pokeapi.co/api/v2/location-area/"

type user struct {
	locations  pokeAPIHelperGo.LocationAreas
	encounters pokeAPIHelperGo.PokeEncounters
	pokedex    map[string]pokeAPIHelperGo.Pokemon
}

func newUser() *user {
	u := user{
		locations:  pokeAPIHelperGo.LocationAreas{Next: base_url},
		encounters: pokeAPIHelperGo.PokeEncounters{},
		pokedex:    make(map[string]pokeAPIHelperGo.Pokemon),
	}
	u.locations.Cache = pokeCache.NewPokeCache(5 * time.Second)
	u.encounters.Cache = pokeCache.NewPokeCache(5 * time.Second)
	return &u
}
