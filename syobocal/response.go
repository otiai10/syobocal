package syobocal

import (
	"encoding/xml"
	"strconv"
	"time"
)

// TitleLookupResponse reporesents xml:TitleLookupResponse.
// ex) http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20150315_000000-
type TitleLookupResponse struct {
	Result struct {
		Code    int    `xml:"Code"`
		Message string `xml:"Message"`
	} `xml:"Result"`
	TitleItems struct {
		Items []TitleItem `xml:"TitleItem"`
	} `xml:"TitleItems"`
}

// TitleItem ...
type TitleItem struct {
	TID           int        `xml:"TID"`           // アニメタイトルID
	LastUpdate    syoboiTime `xml:"LastUpdate"`    // 最近更新された時間
	Title         string     `xml:"Title"`         // アニメタイトル
	ShortTitle    string     `xml:"ShortTitle"`    // タイトルの省略形
	TitleYomi     string     `xml:"TitleYomi"`     // タイトルの日本語読み
	TitleEN       string     `xml:"TitleEN"`       // タイトルを英語にしたらどうなんの的なやつ（英語なければ""）
	TitleFlag     int        `xml:"TitleFlag"`     // たぶん再放送フラグ的なやつ
	Comment       string     `xml:"Comment"`       // コメント（？）なんかいろいろごちゃまぜ
	Category      Category   `xml:"Cat"`           // カテゴリ
	FirstYear     syoboiInt  `xml:"FirstYear"`     // 最初に放送された年
	FirstMonth    syoboiInt  `xml:"FirstMonth"`    // 最初に放送された月
	FirstEndYear  syoboiInt  `xml:"FirstEndYear"`  // 最初の放送が終わった年（未終了なら0）
	FirstEndMonth syoboiInt  `xml:"FirstEndMonth"` // 最初の放送が終わった月（未終了なら0）
	FirstChannel  string     `xml:"FirstCh"`       // 最初に放送された放送局
	Keywords      string     `xml:"Keywords"`      // 検索キーワード. カンマ区切り
	UserPoint     int        `xml:"UserPoint"`     // しょぼかるって投票できるらしいのでそのポイント
	UserPointRank int        `xml:"UserPointRank"` // しょぼかるって投票できるらしいのでその結果ランキング
	SubTitles     string     `xml:"SubTitles"`     // 各放映話のサブタイトル（改行区切り？）
}

type syoboiTime struct{ Value time.Time }

func (t *syoboiTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02 15:04:05"
	var val string
	if err := d.DecodeElement(&val, &start); err != nil {
		return nil
	}
	parsed, err := time.Parse(format, val)
	if err != nil {
		return nil
	}
	*t = syoboiTime{parsed}
	return nil
}

type syoboiInt struct{ Value int }

func (i *syoboiInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var val string
	if err := d.DecodeElement(&val, &start); err != nil {
		return nil
	}
	parsed, err := strconv.Atoi(val)
	if err != nil {
		*i = syoboiInt{0}
		return nil
	}
	*i = syoboiInt{parsed}
	return nil
}
