package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeCache"
)

var registry map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(
		*pokeAPIHelperGo.AreaRespBody,
		*pokeCache.PokeCache,
	) error
}

func commandExit(currMap *pokeAPIHelperGo.AreaRespBody, cache *pokeCache.PokeCache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(currMap *pokeAPIHelperGo.AreaRespBody, cache *pokeCache.PokeCache) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for com, cliCom := range registry {
		fmt.Printf("%v: %v\n", com, cliCom.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(currMap *pokeAPIHelperGo.AreaRespBody, cache *pokeCache.PokeCache) error {
	resp := pokeAPIHelperGo.AreaRespBody{}

	if currMap.Next == "" {
		return fmt.Errorf("no more areas to explore")
	} else if val, exists := cache.Get(currMap.Next); exists {
		if err := json.Unmarshal(val, &resp); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		new_loc, err := pokeAPIHelperGo.ReturnLocations(currMap.Next)
		if err != nil {
			return err
		}
		resp = *new_loc
		jsonData, err2 := json.Marshal(resp)
		if err2 != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err2)
		}
		cache.Add(currMap.Next, jsonData)
	}

	currMap.Next = resp.Next
	currMap.Prev = resp.Prev
	currMap.Result = resp.Result

	for _, area := range currMap.Result {
		fmt.Println(area)
	}
	return nil
}

func commandMapb(currMap *pokeAPIHelperGo.AreaRespBody, cache *pokeCache.PokeCache) error {
	resp := pokeAPIHelperGo.AreaRespBody{}

	if currMap.Prev == "" {
		return fmt.Errorf("no more areas to explore")
	} else if val, exists := cache.Get(currMap.Prev); exists {
		if err := json.Unmarshal(val, &resp); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		new_loc, err := pokeAPIHelperGo.ReturnLocations(currMap.Prev)
		if err != nil {
			return err
		}
		resp = *new_loc
		jsonData, err2 := json.Marshal(resp)
		if err2 != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err2)
		}
		cache.Add(currMap.Prev, jsonData)
	}

	currMap.Next = resp.Next
	currMap.Prev = resp.Prev
	currMap.Result = resp.Result

	for _, area := range currMap.Result {
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
