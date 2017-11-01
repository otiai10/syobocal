package syobocal

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// BaseURL ...
	// http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20171001_000000-
	BaseURL = "http://cal.syoboi.jp/db.php"
)

// Client ...
type Client struct {
	HTTPClient      *http.Client
	BaseURL         string
	LastUpdatedFrom *time.Time
	LastUpdatedTo   *time.Time
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    BaseURL,
	}
}

// Lookup ...
func (c *Client) Lookup() (*TitleLookupResponse, error) {
	if c.LastUpdatedFrom == nil && c.LastUpdatedTo == nil {
		return nil, fmt.Errorf("LastUpdatedの期間指定は必須です。開始日時、終了日時のいずれかを指定する必要があります。")
	}
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}
	req, err := http.NewRequest("GET", c.Build(), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(TitleLookupResponse)
	if err := xml.Unmarshal(b, response); err != nil {
		return nil, err
	}
	if response.Result.Code == 0 {
		xml.Unmarshal(b, &response.Result)
	}
	return response, nil
}

// LastUpdated ...
func (c *Client) LastUpdated(from, to *time.Time) *Client {
	c.LastUpdatedFrom = from
	c.LastUpdatedTo = to
	return c
}

// Build ...
func (c *Client) Build() string {
	return c.BaseURL + "?" + c.Query().Encode()
}

// Query ...
func (c *Client) Query() url.Values {
	v := url.Values{}
	v.Add("Command", "TitleLookup")
	v.Add("TID", "*")
	v.Add("LastUpdate", c.BuildLastUpdated())
	return v
}

// BuildLastUpdated ...
func (c *Client) BuildLastUpdated() string {
	const format = "20060102_150405"
	from := ""
	to := ""
	if c.LastUpdatedFrom != nil {
		from = c.LastUpdatedFrom.Format(format)
	}
	if c.LastUpdatedTo != nil {
		to = c.LastUpdatedTo.Format(format)
	}
	return fmt.Sprintf("%s-%s", from, to)
}
