package command

import (
	"fmt"
	"errors"
)

func requestLocation(url *string, config *Config) error {

	location, err := config.PokeapiClient.GetLocation(
		url,
		config.Pokecache,
		)
	if err != nil {
		return err
	}

	config.LocationNext = location.Next
	config.LocationPrevious = location.Previous
	for _, loc := range location.Results{
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMap(config *Config, params []string) error {

	if len(params) != 0 {
		return errors.New("The map command does not accept parameters. Usage: Pokedex > map")
	}
	return requestLocation(config.LocationNext, config)
}

func commandBmap(config *Config, params []string) error {

	if len(params) != 0 {
		return errors.New("The bmap command does not accept parameters. Usage: Pokedex > bmap")
	}

	if config.LocationPrevious == nil {
		return errors.New("You're on the first page!")
	}
	return requestLocation(config.LocationPrevious, config)
}
