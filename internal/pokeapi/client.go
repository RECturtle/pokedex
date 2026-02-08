package pokeapi

import (
	"net/http"
	"time"

	"github.com/recturtle/pokedexcli/internal/pokecache"
)

const defaultBaseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		httpClient: http.DefaultClient,
		cache:      pokecache.NewCache(20 * time.Second),
	}
}
