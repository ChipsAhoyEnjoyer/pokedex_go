package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	commandRegistry := generateCommandRegistry()

	response := bufio.NewScanner(os.Stdin)
	user := newUser()

	for {
		fmt.Print("Pokedex > ")
		response.Scan()
		if response.Text() == "" {
			fmt.Println("no input...")
			continue
		}
		fullInput := cleanInput(response.Text())
		userCommand, userInput := fullInput[0], fullInput[1]
		if command, ok := commandRegistry[userCommand]; !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(user, userInput)
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
	if len(s) == 1 {
		s = append(s, "")
	}
	return s
}
