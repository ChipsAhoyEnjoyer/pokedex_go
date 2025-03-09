package pokeAPIHelperGo

import "fmt"

type areaRespBody struct {
	Next   string `json:"next"`
	Prev   string `json:"previous"`
	Result []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (c *areaRespBody) GetAreas() {
	for _, area := range c.Result {
		fmt.Println(area)
	}
}
