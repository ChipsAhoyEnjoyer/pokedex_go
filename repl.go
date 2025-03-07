package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	commandRegistry := generateCommandRegistry()

	user := bufio.NewScanner(os.Stdin)
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
			err := command.callback()
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
