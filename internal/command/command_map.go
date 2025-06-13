package command

import (
	"fmt"
	"github.com/rasamadeu/pokedexcli/internal/pokeapi"
)

func requestLocation(url string, config *config) error {

	location, err := pokeapi.GetLocation(url)
	if err != nil {
		return err
	}

	config.next = location.Next
	config.previous = location.Previous
	for _, loc := range location.Results{
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMap(commands map[string]CliCommand, config *config) func() error {
	return func() error {

		url := pokeapi.PokeapiLocationAreaURL
		if config.next != "" {
			url = config.next
		}

		return requestLocation(url, config)
	}
}

func commandBmap(commands map[string]CliCommand, config *config) func() error {
	return func() error {

		if config.previous == "" {
			fmt.Println("You're on the first page!")
			return nil
		}
		url := pokeapi.PokeapiLocationAreaURL

		return requestLocation(url, config)
	}
}
