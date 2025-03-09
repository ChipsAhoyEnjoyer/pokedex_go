package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/poke_api_helper_go"
)

const base_url = "https://pokeapi.co/api/v2/location-area/"

func startRepl() {
	commandRegistry := generateCommandRegistry()

	user := bufio.NewScanner(os.Stdin)
	config := poke_api_helper_go.Config{Next: base_url}

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
			err := command.callback(&config)
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
