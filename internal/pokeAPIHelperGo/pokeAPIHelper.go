package pokeAPIHelperGo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"
)

func FetchData(url string, pokeModel any) error {
	// Accepts a pointer to a model from the models.go file
	// fetches the data from the url and gives it to the model
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeModel)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	return nil
}

func FetchPokemonFromAPI(name string) (*Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name + "/"
	p := &Pokemon{}
	err := FetchData(url, p)
	if err != nil {
		return nil, fmt.Errorf("the pokemon doesn't exist; error finding pokemon ('%v')", name)
	}
	return p, nil
}

func GetLocationData(nextURL string, cache *pokeCache.PokeCache) (*LocationAreas, error) {
	l := &LocationAreas{}
	if nextURL == "" {
		return nil, fmt.Errorf("no more areas to explore")
	} else if val, exists := cache.Get(nextURL); exists {
		err := json.Unmarshal(val, l)
		if err != nil {
			return nil, err
		}
		return l, nil
	}
	err := FetchData(nextURL, l)
	if err != nil {
		return nil, err
	}
	// save to cache
	data, err := json.Marshal(l)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data to add to cache: %v", err)
	}
	cache.Add(nextURL, data)

	return l, nil
}

func GetEncountersData(location string, cache *pokeCache.PokeCache) (*PokeEncounters, error) {
	payload := &PokeEncounters{}

	if location == "" {
		return nil, fmt.Errorf("please type in the name of an area you want to explore")
	} else if val, exists := cache.Get(location); exists {
		if err := json.Unmarshal(val, payload); err != nil {
			return nil, fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
		return payload, nil
	}

	url := "https://pokeapi.co/api/v2/location-area/" + location + "/"
	err := FetchData(url, payload)
	if err != nil {
		return nil, fmt.Errorf("error retreiving encounters; no such area exists ('%v')", location)
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data to add to cache: %v", err)
	}
	cache.Add(location, data)
	return payload, nil

}
