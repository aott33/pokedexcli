package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/aott33/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient 	http.Client
	cache		*pokecache.Cache
}

func NewClient (timeout, interval time.Duration) Client {
	return Client {
		httpClient: http.Client {
			Timeout: 	timeout,
		},
		cache:		pokecache.NewCache(interval),
	}
}

func (c *Client) getJSON(url string) ([]byte, error) {
	var body []byte

	val, ok := c.cache.Get(url)

	if ok {
		body = val	
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return body, err
		}

		defer res.Body.Close()
	
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return body, err
		}

		c.cache.Add(url, body)
	}

	return body, nil

}

func (c *Client) GetLocationAreas(url string) (AreaResponse, error) {
	var areaResp AreaResponse
	
	body, err := c.getJSON(url)
	
	if err != nil {
		return areaResp, err
	}

	err = json.Unmarshal(body, &areaResp)
	if err != nil {
		return areaResp, err
	}

	return areaResp, nil
}

func (c *Client) GetLocationPokemon(url string) (AreaPokemon, error) {
	var areaPoke AreaPokemon
	
	body, err := c.getJSON(url)
	
	if err != nil {
		return areaPoke, err
	}

	err = json.Unmarshal(body, &areaPoke)
	if err != nil {
		return areaPoke, err
	}

	return areaPoke, nil
}
