package command

import (
	"fmt"
	"errors"
)

func requestLocation(config_url *string, config *Config) error {

	location, err := config.PokeapiClient.GetLocation(config_url)
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

func commandMap(config *Config) error {
		return requestLocation(config.LocationNext, config)
}

func commandBmap(config *Config) error {

		if config.LocationPrevious == nil {
			return errors.New("You're on the first page!")
		}
		return requestLocation(config.LocationPrevious, config)
}
