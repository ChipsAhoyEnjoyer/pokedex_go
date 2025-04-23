package pokeAPIHelperGo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"
)

func FetchData(url string, emptyPayload any) (err error) {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&emptyPayload)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	return nil
}

func GetOrCacheLocationData(nextURL string, cache *pokeCache.PokeCache) (locs *LocationAreas, err error) {
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
	err = FetchData(nextURL, l)
	if err != nil {
		return nil, err
	}
	// save to cache
	jsonData, err := json.Marshal(l)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data to add to cache: %v", err)
	}
	cache.Add(nextURL, jsonData)

	return l, nil
}
