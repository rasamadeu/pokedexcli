package command

import "fmt"

// Help command callback function
func commandHelp(commands map[string]CliCommand) func(*Config) error {
	return func(config *Config) error {
		fmt.Printf("\n")
		fmt.Printf("Welcome to the Pokedex!")
		fmt.Printf("Usage:")
		fmt.Printf("\n")
		for _, value := range commands {
			fmt.Printf("%s: %s\n", value.Name, value.Description)
		}
		return nil
	}
}

