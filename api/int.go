package api

import (
	"encoding/xml"
	"strconv"
)

// SyoboiInt ...
type SyoboiInt int

// UnmarshalXML ...
func (i *SyoboiInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var val string
	if err := d.DecodeElement(&val, &start); err != nil {
		return nil
	}
	if parsed, err := strconv.Atoi(val); err != nil {
		*i = 0
	} else {
		*i = SyoboiInt(parsed)
	}
	return nil
}
