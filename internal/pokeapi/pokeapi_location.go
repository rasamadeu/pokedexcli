package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/rasamadeu/pokedexcli/internal/pokecache"
)

const pokeapiLocationAreaURL = pokeapiURL + "/location-area"

type Location struct {
	count    string
	Next     *string
	Previous *string
	Results  []struct{
		         Name string
		         url  string
	         }
}

func (c *Client) GetLocation(url *string, pokecache *pokecache.Cache) (Location, error)  {

	// Check if url corresponds to PokeAPI location endpoint
	reqUrl := pokeapiLocationAreaURL
	if url != nil {
		reqUrl = *url
	}

	// Use cached data if possible.
	// Otherwise, fetch it using http request
	data, ok := pokecache.Get(reqUrl)
	if !ok {
		var err error
		data, err = c.httpGet(reqUrl)
		if err != nil {
			return Location{}, err
		}
	}

	// Add data retrieved to the cache
	pokecache.Add(reqUrl, data)
	var location Location
	err := json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, fmt.Errorf("Error: invalid struct to Unmarshal %s: %v", data, err)
	}

	return location, nil
}
