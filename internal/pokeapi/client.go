package pokeapi

import (
	"net/http"
	"time"

	"github.com/usama-m98/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(10 * time.Second)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
