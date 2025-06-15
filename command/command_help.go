package command

import (
	"errors"
	"fmt"
)

// Help command callback function
func commandHelp(commands map[string]CliCommand) func(*Config, []string) error {
	return func(config *Config, params []string) error {
		if len(params) != 0 {
			return errors.New("The help command does not accept parameters. Usage: Pokedex > help")
		}

		fmt.Printf("\n")
		fmt.Printf("Welcome to the Pokedex!\n\n")
		fmt.Printf("Usage:\n")
		for _, value := range commands {
			fmt.Printf("%s: %s\n", value.Name, value.Description)
		}
		return nil
	}
}

