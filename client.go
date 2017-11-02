package anime

import (
	"net/http"
	"time"

	"github.com/otiai10/anime/syobocal"
)

// Client ...
type Client struct {
	HTTPClient *http.Client
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
	}
}

// Lookup ...
func (c *Client) Lookup() ([]*Anime, error) {
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}
	syoboiclient := syobocal.NewClient()
	syoboiclient.HTTPClient = c.HTTPClient
	from := time.Now().Add(-16 * time.Hour)
	syoboiclient.LastUpdated(&from, nil)
	res, err := syoboiclient.Lookup()
	if err != nil {
		return nil, err
	}
	return CreateAnimeListFromSyobocalResponse(res)
}
