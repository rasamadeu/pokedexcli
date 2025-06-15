package command

import (
	"fmt"
	"errors"
)

func commandExplore(config *Config, params []string) error {

	if len(params) != 1 {
		return errors.New("The explore command needs an area input. Usage: Pokedex > explore <area_name>")
	}

	location := params[0]
	locationPokemons, err := config.PokeapiClient.ExploreLocation(
		location,
		config.Pokecache,
		)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", location, "...")
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationPokemons.PokemonEncounters{
		fmt.Println("-", pokemonEncounter.Pokemon.Name)
	}
	return nil
}
