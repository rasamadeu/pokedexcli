package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/rasamadeu/pokedexcli/internal/pokecache"
)

type LocationPokemons struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocation(area_name string, pokecache *pokecache.Cache) (LocationPokemons, error)  {

	reqUrl := pokeapiLocationAreaURL + "/" + area_name
	// Use cached data if possible.
	// Otherwise, fetch it using http request
	data, ok := pokecache.Get(reqUrl)
	if !ok {
		var err error
		data, err = c.httpGet(reqUrl)
		if err != nil {
			return LocationPokemons{}, err
		}
	}

	// Add data retrieved to the cache
	pokecache.Add(reqUrl, data)
	var pokemons LocationPokemons
	err := json.Unmarshal(data, &pokemons)
	if err != nil {
		return LocationPokemons{}, fmt.Errorf("Error: invalid struct to Unmarshal\ndata: %s\nerr: %v", data, err)
	}

	return pokemons, nil
}
