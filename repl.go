package main

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
)

func startRepl() {
	commandRegistry := generateCommandRegistry()
	user := newUser()

	for {
		response, err := keyEventListener()
		if err != nil {
			fmt.Printf("Error getting key input: %v\n", err)
			continue
		}
		if response == "" {
			fmt.Println("Please enter a command... ")
			continue
		}
		fullInput := cleanInput(response)
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

func keyEventListener() (command string, err error) {
	if err := keyboard.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	typing := true
	input := ""
	prefix := "Pokedex > "

	for typing {
		fmt.Print(prefix, input)
		currLineCharacterCount := len(prefix) + len(input)
		char, key, err := keyboard.GetKey()
		if err != nil {
			return "", err
		}
		switch key {
		case keyboard.KeyEsc:
			return "", nil
		case keyboard.KeyBackspace2:
			if len(input) > 0 {
				input = input[:len(input)-1]
			}
		case keyboard.KeyEnter:
			fmt.Println()
			return input, nil
		default:
			if key == 0 { // check if key pressed is a special character i.e. ENTER, BACKSPACE, CAPS LOCK, etc.
				input += string(char)
			}
		}
		if !typing {
			break
		}
		// clear the current line
		fmt.Print("\r")
		fmt.Print(strings.Repeat(" ", currLineCharacterCount))
		fmt.Print("\r")
	}
	return "", nil
}
