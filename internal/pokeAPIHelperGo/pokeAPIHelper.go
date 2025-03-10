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
