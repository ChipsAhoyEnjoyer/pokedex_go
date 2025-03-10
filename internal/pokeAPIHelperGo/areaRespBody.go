package pokeAPIHelperGo

type AreaRespBody struct {
	Next   string `json:"next"`
	Prev   string `json:"previous"`
	Result []struct {
		Name string `json:"name"`
	} `json:"results"`
}
