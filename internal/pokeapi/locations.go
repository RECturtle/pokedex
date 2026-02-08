package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

const locationAreaPath = "location-area/?offset=0&limit=20"

func (c *Client) InitialLocationAreaURL() string {
	return c.baseURL + locationAreaPath
}

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {
	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		var locationResponse LocationAreaResponse
		if err := json.Unmarshal(data, &locationResponse); err == nil {
			return locationResponse, nil
		}
	}

	// Cache miss — call the API
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationResponse LocationAreaResponse
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationAreaResponse{}, err
	}

	// Store in cache
	c.cache.Add(url, data)

	return locationResponse, nil
}

func (c *Client) GetLocationAreaPokemon(name string) (LocationAreaDetailResponse, error) {
	url := c.baseURL + "location-area/" + name
	// Check cache first
	if data, ok := c.cache.Get(url); ok {
		var locationAreaDetails LocationAreaDetailResponse
		if err := json.Unmarshal(data, &locationAreaDetails); err == nil {
			return locationAreaDetails, nil
		}
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LocationAreaDetailResponse{}, fmt.Errorf("location not found: %s", name)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetailResponse{}, err
	}

	var locationAreaDetails LocationAreaDetailResponse
	if err := json.Unmarshal(data, &locationAreaDetails); err != nil {
		return LocationAreaDetailResponse{}, err
	}
	c.cache.Add(url, data)
	return locationAreaDetails, nil
}
