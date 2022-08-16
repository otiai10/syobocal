package api

import (
	"encoding/xml"
	"time"
)

const (
	queryTimeFormat = "20060102_150405"
)

// SyoboiTime ...
type SyoboiTime time.Time

// UnmarshalXML ...
func (t *SyoboiTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02 15:04:05"
	var val string
	if err := d.DecodeElement(&val, &start); err != nil {
		return nil
	}
	if parsed, err := time.Parse(format, val); err == nil {
		*t = SyoboiTime(parsed)
	}
	return nil
}
