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

func (c *Client) GetLocationAreas(url string) (AreaResponse, error) {
	var areaResp AreaResponse
	var body []byte

	val, ok := c.cache.Get(url)

	if ok {
		body = val	
	} else {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return areaResp, err
		}

		defer res.Body.Close()
	
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return areaResp, err
		}

		c.cache.Add(url, body)
	}
	
	err := json.Unmarshal(body, &areaResp)
	if err != nil {
		return areaResp, err
	}

	return areaResp, nil
}
