package main

import (
	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
)

func registryMiddleware(f func(*user, string, map[string]cliCommand) error) func(*user, string) error {
	return func(u *user, s string) error {
		commands := generateCommands() // TODO: Don't want to regenerate the commands list everytime the user calls help
		if err := f(u, s, commands); err != nil {
			return err
		}
		return nil
	}
}

func mapMiddleware(f func(*user, string, *pokeAPIHelperGo.LocationAreas) error) func(*user, string) error {
	return func(u *user, s string) error {
		locations, err := pokeAPIHelperGo.GetLocationData(u.locations.Next, u.locations.Cache)
		if err != nil {
			return err
		}
		return f(u, s, locations)
	}
}

func mapbMiddleware(f func(*user, string, *pokeAPIHelperGo.LocationAreas) error) func(*user, string) error {
	return func(u *user, s string) error {
		locations, err := pokeAPIHelperGo.GetLocationData(u.locations.Prev, u.locations.Cache)
		if err != nil {
			return err
		}
		return f(u, s, locations)
	}
}

func exploreMiddleware(f func(*user, string, *pokeAPIHelperGo.PokeEncounters) error) func(*user, string) error {
	return func(u *user, s string) error {
		encounters, err := pokeAPIHelperGo.GetEncountersData(s, u.encounters.Cache)
		if err != nil {
			return err
		}
		return f(u, s, encounters)
	}
}

func catchMiddleware(f func(*user, string, *pokeAPIHelperGo.Pokemon) error) func(*user, string) error {
	return func(u *user, s string) error {
		p, err := pokeAPIHelperGo.FetchPokemonFromAPI(s)
		if err != nil {
			return err
		}
		return f(u, s, p)
	}
}
