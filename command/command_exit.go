package command

import (
	"errors"
	"fmt"
	"os"
)

// Exit command callback function
func commandExit(config *Config, params []string) error {

	if len(params) != 0 {
		return errors.New("The exit command does not accept parameters. Usage: Pokedex > exit")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	config.Pokecache.Kill()
	os.Exit(0)
	return nil
}
