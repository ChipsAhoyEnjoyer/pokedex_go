package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
)

var registry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeAPIHelperGo.LocationAreas) error
}

func commandExit(locations *pokeAPIHelperGo.LocationAreas) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(locations *pokeAPIHelperGo.LocationAreas) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for com, cliCom := range registry {
		fmt.Printf("%v: %v\n", com, cliCom.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(locations *pokeAPIHelperGo.LocationAreas) error {
	resp := pokeAPIHelperGo.LocationAreas{}

	if locations.Next == "" {
		return fmt.Errorf("no more areas to explore")
	} else if val, exists := locations.Cache.Get(locations.Next); exists {
		if err := json.Unmarshal(val, &resp); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		new_loc, err := pokeAPIHelperGo.ReturnLocations(locations.Next)
		if err != nil {
			return err
		}
		resp = *new_loc
		jsonData, err2 := json.Marshal(resp)
		if err2 != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err2)
		}
		locations.Cache.Add(locations.Next, jsonData)
	}

	locations.Next = resp.Next
	locations.Prev = resp.Prev
	locations.Result = resp.Result

	for _, area := range locations.Result {
		fmt.Println(area)
	}
	return nil
}

func commandMapb(locations *pokeAPIHelperGo.LocationAreas) error {
	resp := pokeAPIHelperGo.LocationAreas{}

	if locations.Prev == "" {
		return fmt.Errorf("no more areas to explore")
	} else if val, exists := locations.Cache.Get(locations.Prev); exists {
		if err := json.Unmarshal(val, &resp); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		new_loc, err := pokeAPIHelperGo.ReturnLocations(locations.Prev)
		if err != nil {
			return err
		}
		resp = *new_loc
		jsonData, err2 := json.Marshal(resp)
		if err2 != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err2)
		}
		locations.Cache.Add(locations.Prev, jsonData)
	}

	locations.Next = resp.Next
	locations.Prev = resp.Prev
	locations.Result = resp.Result

	for _, area := range locations.Result {
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
