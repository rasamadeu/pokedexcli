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
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

// This function return a map containing all possible Cli commands
func CreateCommands() map[string]CliCommand {

	commands := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}

	commands["map"] = CliCommand{
		Name:        "map",
		Description: "Displays the next 20 locations",
		Callback:    commandMap,
	}

	commands["bmap"] = CliCommand{
		Name:        "bmap",
		Description: "Displays the previous 20 locations",
		Callback:    commandBmap,
	}

	commands["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp(commands),
	}

	return commands
}
