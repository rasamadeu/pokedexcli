package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const PokeapiURL = "https://pokeapi.co/api/v2"
const PokeapiLocationAreaURL = PokeapiURL + "/location-area"

type Location struct {
	count    string
	Next     string
	Previous string
	Results  []struct{
		         Name string
		         url  string
	         }
}

func GetLocation(url string) (Location, error)  {

	var location Location
	// Check if url corresponds to PokeAPI location endpoint
	if (!strings.HasPrefix(url, PokeapiLocationAreaURL)) {
		return location, fmt.Errorf("Error: %s is not a valid Pokeapi location endpoint", url)
	}

	// Retrieve data from url
	res, err := http.Get(url)
	if err != nil {
		return location, fmt.Errorf("Error: failed GET %s: %v", url, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return location, fmt.Errorf("Error: failed reading data from response to %s: %v", url, err)
	}

	err = json.Unmarshal(data, &location)
	if err != nil {
		return location, fmt.Errorf("Error: invalid struct to Unmarshal %s: %v", data, err)
	}

	return location, nil
}
