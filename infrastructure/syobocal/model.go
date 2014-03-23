package syobocal

import "encoding/xml"

type Result struct {
	Code    string
	Message string
}
type TitleItem struct {
	XMLName    xml.Name `xml:"TitleItem"`
	TID        string   `xml:"TID"`
	LastUpdate string   `xml:"LastUpdate"`
	Title      string   `xml:"Title"`
	ShortTitle string   `xml:"ShortTitle"`
	TitleYomi  string   `xml:"TitleYomi"`
	Comment    string   `xml:"Comment"`
	FirstYear  int      `xml:"FirstYear"`
	FirstMonth int      `xml:"FirstMonth"`
	Cat        int      `xml:"Cat"`
}
type TitleItems struct {
	XMLName   xml.Name `xml:"TitleItems"`
	TitleItem []TitleItem
}
type Response struct {
	XMLName    xml.Name   `xml:"TitleLookupResponse"`
	Result     Result     `xml:"Result"`
	TitleItems TitleItems `xml:"TitleItems"`
}
