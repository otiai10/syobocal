package syobocal

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	// APIEndpoint しょぼかるのエンドポイント
	// https://sites.google.com/site/syobocal/spec/db-php
	APIEndpoint = "http://cal.syoboi.jp/db.php"
	// CommandTitleLookup タイトルデータ
	// https://sites.google.com/site/syobocal/spec/db-php#TOC-TitleLookup-
	CommandTitleLookup = "TitleLookup"
	// SyoboiTimeFormat しょぼかるの時間フォーマット
	SyoboiTimeFormat = "20060102_150405"
)

// Category しょぼかる的カテゴリ
// 1 アニメ
// 10 アニメ(終了/再放送)
// 7 OVA
// 5 アニメ関連
// 4 特撮
// 8 映画
// 3 テレビ
// 2 ラジオ
// 6 メモ
// 0 その他
type Category int

// Client is client API for syobocal.
// @see https://sites.google.com/site/syobocal/spec/db-php
type Client struct {
	Endpoint string
	Service  Service
	from     time.Time
	to       time.Time
	Query    url.Values
}

// NewClient returns Client.
func NewClient(ss ...Service) *Client {
	client := &Client{
		Service:  new(defaultService),
		Endpoint: APIEndpoint,
		Query: url.Values{
			"TID": []string{"*"},
		},
	}
	if len(ss) != 0 {
		client.Service = ss[0]
	}
	return client
}

// TitleLookup Commandのセッター
func (client *Client) TitleLookup() *Client {
	client.Query.Add("Command", CommandTitleLookup)
	return client
}

// From LastUpdateのセッター
func (client *Client) From(t time.Time) *Client {
	client.from = t
	return client
}

// To LastUpdateのセッター
func (client *Client) To(t time.Time) *Client {
	client.to = t
	return client
}

// LastUpdate LastUpdateのセッター
func (client *Client) LastUpdate(from, to time.Time) *Client {
	client.From(from).To(to)
	return client
}

func (client *Client) getURL() string {
	var f, t string
	if !client.from.IsZero() {
		f = client.from.Format(SyoboiTimeFormat)
	}
	if !client.to.IsZero() {
		t = client.to.Format(SyoboiTimeFormat)
	}
	if len(f)+len(t) != 0 {
		client.Query.Add("LastUpdate", fmt.Sprintf("%s-%s", f, t))
	}
	u := url.URL{
		Path:     client.Endpoint,
		RawQuery: client.Query.Encode(),
	}
	return u.String()
}

// Do ...
func (client *Client) Do(term ...time.Time) (TitleLookupResponse, error) {
	res, err := http.Get(client.getURL())
	if err != nil {
		return TitleLookupResponse{}, err
	}
	tlr := TitleLookupResponse{}
	decoder := xml.NewDecoder(res.Body)
	if err := decoder.Decode(&tlr); err != nil {
		return tlr, err
	}
	return tlr, nil
}
