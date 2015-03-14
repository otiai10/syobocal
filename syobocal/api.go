package syobocal

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
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
type Client struct{}

// NewClient returns Client.
func NewClient() *Client {
	return &Client{}
}

// TitleLookup ...
// ex) http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20150315_000000-
func (client *Client) TitleLookup(from, to time.Time) (TitleLookupResponse, error) {
	res, err := http.Get("http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20150315_000000-")
	if err != nil {
		return TitleLookupResponse{}, err
	}
	fmt.Println(res, err)

	tlr := TitleLookupResponse{}
	decoder := xml.NewDecoder(res.Body)
	if err := decoder.Decode(&tlr); err != nil {
		return tlr, err
	}
	return tlr, nil
}
