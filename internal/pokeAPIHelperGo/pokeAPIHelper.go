package pokeAPIHelperGo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnLocations(url string) (*LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	var c LocationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return &c, nil
}

func ReturnPokeEncounters(url string) (*PokeEncounters, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	var e PokeEncounters
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&e)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return &e, nil
}
func ReturnPokemon(url string) (*Pokemon, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	var p Pokemon
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&p)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return &p, nil
}
