package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const pokeapiURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient *http.Client

}

func NewClient(timeout time.Duration) *Client {
	return &Client {
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: failed GET %s:\n%v", url, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: failed reading data from response to %s:\n%v", url, err)
	}
	return data, nil
}
