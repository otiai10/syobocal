package syobocal

import (
	"encoding/xml"
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
	TID        string     `xml:"TID"`
	LastUpdate syoboiTime `xml:"LastUpdate"`
}

type syoboiTime struct{ time.Time }

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
