package pokeapi

import (
	"net/http"
	"time"

	"github.com/recturtle/pokedexcli/internal/pokecache"
)

const DefaultBaseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		cache:      pokecache.NewCache(20 * time.Second),
	}
}
