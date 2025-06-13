package command

import "fmt"

// Help command callback function
func commandHelp(commands map[string]CliCommand) func() error {
	return func() error {
		fmt.Println("\nWelcome to the Pokedex!")
		fmt.Println("Usage:\n")
		for _, value := range commands {
			fmt.Printf("%s: %s\n", value.Name, value.Description)
		}
		return nil
	}
}

