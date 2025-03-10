package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"
)

const base_url = "https://pokeapi.co/api/v2/location-area/"

func startRepl() {
	commandRegistry := generateCommandRegistry()

	user := bufio.NewScanner(os.Stdin)
	resp := &pokeAPIHelperGo.LocationAreas{Next: base_url}
	resp.Cache = pokeCache.NewPokeCache(5 * time.Second)

	for {
		fmt.Print("Pokedex > ")
		user.Scan()
		userInput := cleanInput(user.Text())[0]
		if len(userInput) == 0 {
			continue
		}
		if command, ok := commandRegistry[userInput]; !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(resp)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	t := strings.ToLower(text)
	s := strings.Fields(t)
	return s
}
