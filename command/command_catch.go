package command

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(config *Config, params []string) error {

	if len(params) != 1 {
		return errors.New("The catch command needs a Pokemon name. Usage: Pokedex > catch <pokemon>")
	}

	pokemonName := params[0]
	pokemon, err := config.PokeapiClient.GetPokemon(
		pokemonName,
		config.Pokecache,
		)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	
	// The lower the pokemon base experience, the higher the likelihood of catching it
	maxBaseExperience := 700
	if rand.Int() % maxBaseExperience < pokemon.BaseExperience {
		fmt.Println(pokemon.Name, "escaped!")
		return nil
	}
	fmt.Println(pokemon.Name, "was caught!")
	config.Pokedex[pokemon.Name] = pokemon
	return nil
}
