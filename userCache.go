package main

import (
	"time"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"
)

const base_url = "https://pokeapi.co/api/v2/location-area/"

type userCache struct {
	locations  pokeAPIHelperGo.LocationAreas
	encounters pokeAPIHelperGo.PokeEncounters
}

func newUser() *userCache {
	u := userCache{
		locations:  pokeAPIHelperGo.LocationAreas{Next: base_url},
		encounters: pokeAPIHelperGo.PokeEncounters{},
	}
	u.locations.Cache = pokeCache.NewPokeCache(5 * time.Second)
	u.encounters.Cache = pokeCache.NewPokeCache(5 * time.Second)
	return &u
}
