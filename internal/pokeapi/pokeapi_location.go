package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client) GetLocation(url *string) (Location, error)  {

	// Check if url corresponds to PokeAPI location endpoint
	reqUrl := pokeapiLocationAreaURL
	if url != nil {
		reqUrl = *url
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return Location{}, err
	}

	// Retrieve data from url
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, fmt.Errorf("Error: failed GET %s: %v", reqUrl, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("Error: failed reading data from response to %s: %v", reqUrl, err)
	}

	var location Location
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, fmt.Errorf("Error: invalid struct to Unmarshal %s: %v", data, err)
	}

	return location, nil
}
