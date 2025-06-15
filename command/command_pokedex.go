package command

import (
	"fmt"
	"errors"
)

func commandPokedex(config *Config, params []string) error {

	if len(params) != 0 {
		return errors.New("The pokedex command does not accept parameters. Usage: Pokedex > pokedex")
	}

	fmt.Println("Your Pokedex:")
	for pokemon, _ := range config.Pokedex {
		fmt.Printf("  - %s\n", pokemon)
	}
	return nil
}

