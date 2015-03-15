package factory

import (
	"fmt"
	"testing"
	"time"

	"github.com/otiai10/animapi/syobocal"
	. "github.com/otiai10/mint"
)

func TestParseRawComment(t *testing.T) {
	Expect(t, true).ToBe(true)
	res, err := syobocal.NewClient().TitleLookup(time.Now(), time.Now())
	Expect(t, err).ToBe(nil)

	songs := parseRawComment(res.TitleItems.Items[0].TID, res.TitleItems.Items[0].Comment)
	fmt.Printf("-----\n%+v\n", songs)
}
