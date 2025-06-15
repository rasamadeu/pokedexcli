package command

import (
	"github.com/rasamadeu/pokedexcli/internal/pokeapi"
	"github.com/rasamadeu/pokedexcli/internal/pokecache"
)

type Config struct {
	PokeapiClient    *pokeapi.Client
	LocationNext     *string
	LocationPrevious *string
	Pokecache        *pokecache.Cache
	Pokedex          map[string]pokeapi.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, []string) error
}

// This function return a map containing all possible Cli commands
func CreateCommands() map[string]CliCommand {

	commands := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex. Usage: exit",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 locations. Usage: map",
			Callback:    commandMap,
		},
		"bmap": {
			Name:        "bmap",
			Description: "Displays the previous 20 locations. Usage: bmap",
			Callback:    commandBmap,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the Pokemon found in a location. Usage: explore <location-area>",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a Pokemon. Usage: catch <pokemon>",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Show details of a caught Pokemon. Usage: catch <pokemon>",
			Callback:    commandInspect,
		},
	}


	commands["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message. Usage: help",
		Callback:    commandHelp(commands),
	}

	return commands
}
