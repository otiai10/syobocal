package infrastructure

import "encoding/xml"

// とりあえずしょぼCalからくる
// レスポンスのすべてのフィールドを
// Unmarshalして拾います
type SyobocalResponse struct {
	XMLName    xml.Name     `xml:"TitleLookupResponse"`
	Result     LookupResult `xml:"Result"`
	TitleItems TitleItems   `xml:"TitleItems"`
}
type LookupResult struct {
	Code    int    `xml:"Code"`
	Message string `xml:"Message"`
}
type TitleItems struct {
	Items []TitleItem `xml:"TitleItem"`
}
type TitleItem struct {
	TID          int    `xml:"TID"`
	LastUpdate   string `xml:"LastUpdate"`
	Title        string `xml:"Title"`
	ShortTitle   string `xml:"ShortTitle"`
	TitleYomi    string `xml:"TitleYomi"`
	TitleEnglish string `xml:"TitleEN"`
	// TODO: 文字列処理してstruct化
	Comment       string `xml:"Comment"` //とりあえず
	Category      int    `xml:"Cat"`
	Flag          string `xml:"TitleFlag"`
	FirstYear     string `xml:"FirstYear"`
	FirstMonth    string `xml:"FirstMonth"`
	FirstEndYear  string `xml:"FirstEndYear"`
	FirstEndMonth string `xml:"FirstEndMonth"`
	FirstChannel  string `xml:"FirstCh"`
	Keywords      string `xml:"Keywords"`
	UserPoint     string `xml:"UserPoint"`
	UserPointRank string `xml:"UserPointRank"`
	// TODO: 文字列処理してstruct化
	SubTitles string `xml:"SubTitles"` //とりあえず
}

func ConvertBytes2Response(bytes []byte) (res SyobocalResponse, e error) {
	e = xml.Unmarshal(bytes, &res)
	return
}
