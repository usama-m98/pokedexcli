package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(endpoint *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if endpoint != nil {
		url = *endpoint
	}

	if dat, ok := c.cache.Get(url); ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			log.Fatal("failed to decode data")
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		log.Fatal("failed to decode data")
	}

	c.cache.Add(url, dat)

	return locationArea, nil
}
