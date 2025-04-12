package main

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
