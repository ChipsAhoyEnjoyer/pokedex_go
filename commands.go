package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func commandCatch(userData *userCache, input string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + input + "/"
	pokemon, err := pokeAPIHelperGo.ReturnPokemon(url)
	if err != nil {
		return fmt.Errorf("the pokemon doesn't exist; error finding pokemon ('%v'): %v", input, err)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", input)
	playerRoll := rand.Intn(pokemon.BaseExperience)
	catchRate := rand.Intn(pokemon.BaseExperience)
	if pokemon.BaseExperience > 200 {
		playerRoll = (playerRoll / 2) + 50
	} else if pokemon.BaseExperience < 100 {
		catchRate /= 2
	}
	fmt.Printf("player : %v\n", playerRoll)
	fmt.Printf("roll : %v\n", catchRate)
	catch := playerRoll >= catchRate
	if catch {
		fmt.Printf("%v was caught!\n", input)
		userData.pokedex[input] = *pokemon
	} else {
		fmt.Printf("%v escaped...\n", input)
	}
	return nil
}
func commandInspect(userData *userCache, input string) error {
	if input == "" {
		return fmt.Errorf("please enter a Pokemon name after the 'inspect' command")
	}
	pokemonInfo, exists := userData.pokedex[input]
	if !exists {
		fmt.Printf("%v has not been caught yet\n", input)
		return nil
	}
	fmt.Printf(`
	Name : %v
	Base Experience: %v
	Height : %v
	Weight : %v
	`,
		pokemonInfo.Name,
		pokemonInfo.BaseExperience,
		pokemonInfo.Height,
		pokemonInfo.Weight,
	)
	fmt.Println("\nTypes:")
	for _, t := range pokemonInfo.Types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}
	fmt.Println("\nAbilities:")
	for _, a := range pokemonInfo.Abilities {
		fmt.Printf(" - %v\n", a.Ability.Name)
	}
	fmt.Println("\nStats:")
	for _, s := range pokemonInfo.Stats {
		fmt.Printf("%v: %v\n", s.Stat.Name, s.BaseStat)
	}
	return nil
}
func commandPokedex(userData *userCache, input string) error {
	fmt.Println("Your Pokedex:")
	for k := range userData.pokedex {
		fmt.Printf(" - %v\n", k)
	}
	return nil
}

func generateCommandRegistry() map[string]cliCommand {
	registry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex\n",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message\n",
			callback:    commandHelp,
		},
		"map": {
			name: "map",
			description: `Displays the names of the next 20 location areas in the Pokemon world. 
Each subsequent call to map should display the next 20 locations, 
and so on
`,
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: `Displays the names of the previous 20 location areas in the Pokemon world. 
Each subsequent call to map should display the previous 20 locations, 
and so on
`,
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: `It takes the name of a location area as an argument and returns a list of
pokemon in that area.
`,
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: `Throw a Pokeball at a Pokemon for a chance to capture it.
The stronger the Pokemon, the harder it is to catch!
`,
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: `Inspect your Pokemon.
Enter your caught Pokemon's name as an argument to view its info`,
			callback: commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "`View your Pokedex.",
			callback:    commandPokedex,
		},
	}
	return registry
}
