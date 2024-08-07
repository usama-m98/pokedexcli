package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonsInArea(name string) (LocationPokemons, error) {
	url := baseURL + "/location-area/" + name

	if poke, ok := c.cache.Get(url); ok {
		pokemons := LocationPokemons{}
		err := json.Unmarshal(poke, &pokemons)
		if err != nil {
			return pokemons, err
		}
		return pokemons, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationPokemons{}, err
	}

	pokemons := LocationPokemons{}
	err = json.Unmarshal(dat, &pokemons)
	if err != nil {
		return pokemons, err
	}

	c.cache.Add(url, dat)

	return pokemons, nil
}
