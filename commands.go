package main

import (
	"fmt"
	"os"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/poke_api_helper_go"
)

var registry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*poke_api_helper_go.Config) error
}

func commandExit(c *poke_api_helper_go.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *poke_api_helper_go.Config) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for com, cliCom := range registry {
		fmt.Printf("%v: %v\n", com, cliCom.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(c *poke_api_helper_go.Config) error {
	if c.Next == "" {
		return fmt.Errorf("no more areas to explore")
	}
	new_config, err := poke_api_helper_go.ReturnLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = new_config.Next
	c.Prev = new_config.Prev
	c.Result = new_config.Result

	areas := c.GetAreas()
	for _, area := range areas {
		fmt.Println(area)
	}
	return nil
}

func commandMapb(c *poke_api_helper_go.Config) error {
	if c.Prev == "" {
		return fmt.Errorf("no more areas to explore")
	}
	new_config, err := poke_api_helper_go.ReturnLocations(c.Prev)
	if err != nil {
		return err
	}

	c.Next = new_config.Next
	c.Prev = new_config.Prev
	c.Result = new_config.Result

	areas := c.GetAreas()
	for _, area := range areas {
		fmt.Println(area)
	}
	return nil
}

func generateCommandRegistry() map[string]cliCommand {
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: `displays the names of the next 20 location areas in the Pokemon world. 
			              Each subsequent call to map should display the next 20 locations, 
						  and so on`,
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: `displays the names of the previous 20 location areas in the Pokemon world. 
			              Each subsequent call to map should display the previous 20 locations, 
						  and so on`,
			callback: commandMapb,
		},
	}
	return registry
}
