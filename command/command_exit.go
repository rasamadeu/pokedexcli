package command

import (
	"fmt"
	"os"
)

// Exit command callback function
func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	config.Pokecache.Kill()
	os.Exit(0)
	return nil
}
