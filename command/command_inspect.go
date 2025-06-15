package command

import (
	"fmt"
	"errors"
)

func commandInspect(config *Config, params []string) error {

	if len(params) != 1 {
		return errors.New("The inspect command needs a Pokemon name. Usage: Pokedex > inspect <pokemon>")
	}

	pokemonName := params[0]
	pokemon, ok := config.Pokedex[pokemonName]
	if !ok {
		return errors.New("You can only inspect Pokemons you have caught!\nTry to catch it first!")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range pokemon.Types {
		fmt.Printf("  -%s\n", stat.Type.Name)
	}
	return nil
}
