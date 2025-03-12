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
	callback    func(*userCache, string) error
}

func commandExit(userData *userCache, input string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(userData *userCache, input string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for com, cliCom := range registry {
		fmt.Printf("%v: %v\n", com, cliCom.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(userData *userCache, input string) error {
	resp := pokeAPIHelperGo.LocationAreas{}

	if userData.locations.Next == "" {
		return fmt.Errorf("no more areas to explore")
	} else if val, exists := userData.locations.Cache.Get(userData.locations.Next); exists {
		if err := json.Unmarshal(val, &resp); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		new_loc, err := pokeAPIHelperGo.ReturnLocations(userData.locations.Next)
		if err != nil {
			return err
		}
		resp = *new_loc
		jsonData, err := json.Marshal(resp)
		if err != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err)
		}
		userData.locations.Cache.Add(userData.locations.Next, jsonData)
	}

	userData.locations.Next = resp.Next
	userData.locations.Prev = resp.Prev
	userData.locations.Result = resp.Result

	for _, area := range userData.locations.Result {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(userData *userCache, input string) error {
	resp := pokeAPIHelperGo.LocationAreas{}

	if userData.locations.Prev == "" {
		return fmt.Errorf("no more areas to explore")
	} else if val, exists := userData.locations.Cache.Get(userData.locations.Prev); exists {
		if err := json.Unmarshal(val, &resp); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		new_loc, err := pokeAPIHelperGo.ReturnLocations(userData.locations.Prev)
		if err != nil {
			return err
		}
		resp = *new_loc
		jsonData, err := json.Marshal(resp)
		if err != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err)
		}
		userData.locations.Cache.Add(userData.locations.Prev, jsonData)
	}

	userData.locations.Next = resp.Next
	userData.locations.Prev = resp.Prev
	userData.locations.Result = resp.Result

	for _, area := range userData.locations.Result {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(userData *userCache, input string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + input + "/"

	data := &pokeAPIHelperGo.PokeEncounters{}
	if input == "" {
		return fmt.Errorf("please type in the name of an area you want to explore")
	} else if val, exists := userData.encounters.Cache.Get(input); exists {
		if err := json.Unmarshal(val, data); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		encounters, err := pokeAPIHelperGo.ReturnPokeEncounters(url)
		if err != nil {
			return fmt.Errorf("error retreiving encounters; no such area exists ('%v')", input)
		}
		data = encounters
		jsonData, err := json.Marshal(encounters)
		if err != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err)
		}
		userData.encounters.Cache.Add(input, jsonData)
	}
	fmt.Printf("Exploring %v...\n", input)
	fmt.Println("Found Pokemon:")
	for _, mon := range data.Encounters {
		fmt.Printf("- %v\n", mon.Pokemon.Name)
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
			description: `Displays the names of the next 20 location areas in the Pokemon world. 
			              Each subsequent call to map should display the next 20 locations, 
						  and so on`,
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: `Displays the names of the previous 20 location areas in the Pokemon world. 
			              Each subsequent call to map should display the previous 20 locations, 
						  and so on`,
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: `It takes the name of a location area as an argument and returns a list of
						  pokemon in that area.`,
			callback: commandExplore,
		},
	}
	return registry
}
