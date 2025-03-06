package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	user := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		user.Scan()
		userInput := cleanInput(user.Text())
		if len(userInput) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", userInput[0])
	}
}

func cleanInput(text string) []string {
	t := strings.ToLower(text)
	s := strings.Fields(t)
	return s
}
