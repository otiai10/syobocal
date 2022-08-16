package syobocal

// // Client ...
// type Client struct {
// 	HTTPClient *http.Client
// 	Verbose    bool
// }

// // NewClient ...
// func NewClient() *Client {
// 	return &Client{
// 		HTTPClient: http.DefaultClient,
// 	}
// }

// // Lookup ...
// func (c *Client) Lookup(term ...time.Time) ([]*Anime, error) {
// 	from := time.Now().Add(-16 * time.Hour)
// 	if len(term) > 0 {
// 		from = term[0]
// 	}
// 	var to *time.Time
// 	if len(term) > 1 && term[1].Before(from) {
// 		to = &term[1]
// 	}

// 	if c.HTTPClient == nil {
// 		c.HTTPClient = http.DefaultClient
// 	}
// 	syoboiclient := syobocal.NewClient()
// 	syoboiclient.HTTPClient = c.HTTPClient
// 	syoboiclient.LastUpdated(&from, to)
// 	syoboiclient.Verbose = c.Verbose
// 	res, err := syoboiclient.Lookup()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return CreateAnimeListFromSyobocalResponse(res)
// }
