package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/rasamadeu/pokedexcli/internal/pokecache"
)

const pokeapiPokemonURL = pokeapiURL + "/pokemon"

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort    int `json:"effort"`
		Stat      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}`json:"stat"`
	}`json:"stats"`
	Types          []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
	}`json:"types"`
}

func (c *Client) GetPokemon(pokemonName string, pokecache *pokecache.Cache) (Pokemon, error)  {

	reqUrl := pokeapiPokemonURL + "/" + pokemonName
	// Use cached data if possible.
	// Otherwise, fetch it using http request
	data, ok := pokecache.Get(reqUrl)
	if !ok {
		var err error
		data, err = c.httpGet(reqUrl)
		if err != nil {
			return Pokemon{}, err
		}
	}

	// Add data retrieved to the cache
	pokecache.Add(reqUrl, data)
	var pokemon Pokemon
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error: Pokemon not found!\ndata: %s\nerr: %v", data, err)
	}

	return pokemon, nil
}
