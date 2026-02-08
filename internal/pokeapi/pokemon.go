package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
)

func (c *Client) CatchPokemon(name string) (Pokemon, error) {
	url := c.baseURL + "pokemon/" + name
	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Pokemon{}, fmt.Errorf("Pokemon not found: %s", name)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	if !c.catchCalculation(pokemon.BaseExperience) {
		return pokemon, fmt.Errorf("%s escaped!", pokemon.Name)
	}

	return pokemon, nil
}

func (c *Client) catchCalculation(baseExp int) bool {
	const threshold = 50
	chance := rand.Intn(baseExp)
	return chance < threshold
}
