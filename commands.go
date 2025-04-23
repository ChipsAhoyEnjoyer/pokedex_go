package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/ChipsAhoyEnjoyer/pokedex_go/internal/pokeAPIHelperGo"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*user, string) error
}

// TODO: Simulate Pokemon battles
// TODO: Change Pokemon descriptions to have levels instead of base experience
// TODO: Add tests for commands
// TODO: Add a saved state for users
// TODO: Cache the pokedex info for mons instead of saving all the info to mem

func commandExit(user *user, input string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(user *user, input string, locations *pokeAPIHelperGo.LocationAreas) error {
	user.locations.Next = locations.Next
	user.locations.Prev = locations.Prev
	user.locations.Result = locations.Result
	for _, area := range user.locations.Result {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(user *user, input string, locations *pokeAPIHelperGo.LocationAreas) error {
	user.locations.Next = locations.Next
	user.locations.Prev = locations.Prev
	user.locations.Result = locations.Result
	for _, area := range user.locations.Result {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(user *user, input string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + input + "/"

	data := &pokeAPIHelperGo.PokeEncounters{}
	if input == "" {
		return fmt.Errorf("please type in the name of an area you want to explore")
	} else if val, exists := user.encounters.Cache.Get(input); exists {
		if err := json.Unmarshal(val, data); err != nil {
			return fmt.Errorf("error unmarshaling data from cache: '%v'", err)
		}
	} else {
		encounters, err := pokeAPIHelperGo.FetchData(url, &pokeAPIHelperGo.PokeEncounters{})
		if err != nil {
			return fmt.Errorf("error retreiving encounters; no such area exists ('%v')", input)
		}
		data = encounters
		jsonData, err := json.Marshal(encounters)
		if err != nil {
			return fmt.Errorf("error marshaling data to add to cache: %v", err)
		}
		user.encounters.Cache.Add(input, jsonData)
	}
	fmt.Printf("Exploring %v...\n", input)
	fmt.Println("Found Pokemon:")
	for _, mon := range data.Encounters {
		fmt.Printf("- %v\n", mon.Pokemon.Name)
	}

	return nil
}

func commandCatch(user *user, input string) error {
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
		user.pokedex[input] = *pokemon
	} else {
		fmt.Printf("%v escaped...\n", input)
	}
	return nil
}

func commandInspect(user *user, input string) error {
	if input == "" {
		return fmt.Errorf("please enter a Pokemon name after the 'inspect' command")
	}
	pokemonInfo, exists := user.pokedex[input]
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

func commandPokedex(user *user, input string) error {
	fmt.Println("Your Pokedex:")
	for k := range user.pokedex {
		fmt.Printf(" - %v\n", k)
	}
	return nil
}

func commandHelp(user *user, input string, registry map[string]cliCommand) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for com, cliCom := range registry {
		fmt.Printf("%v: %v\n", com, cliCom.description)
	}
	fmt.Println("")
	return nil
}

func addCommand(
	registry map[string]cliCommand,
	command string,
	description string,
	callback func(*user, string) error,
) {
	registry[command] = cliCommand{
		name:        command,
		description: description,
		callback:    callback,
	}
}

func generateCommands() map[string]cliCommand {
	registry := make(map[string]cliCommand)

	addCommand(registry, "exit", "Exit the Pokedex.", commandExit)
	addCommand(registry, "help", "Displays a help message.", registryMiddleware(commandHelp))
	addCommand(registry, "map", "Displays the names of the next 20 location areas in the Pokemon world.", commandMap)
	addCommand(registry, "mapb", "Displays the names of the previous 20 location areas in the Pokemon world.", commandMapb)
	addCommand(registry, "explore", "Takes a location area as an argument and lists all Pokemon in that area.", commandExplore)
	addCommand(registry, "catch", "Throw a Pokeball for a chance to capture it a Pokemon.", commandCatch)
	addCommand(registry, "inspect", "Inspect your Pokemon.", commandInspect)
	addCommand(registry, "pokedex", "View your Pokedex.", commandPokedex)

	return registry
}
