package factory

import (
	"encoding/xml"
	"log"
	"os"
	"testing"

	"github.com/otiai10/animapi/syobocal"
	. "github.com/otiai10/mint"
)

const SHIROBAKO = 3524

func TestParseRawComment(t *testing.T) {

	res := fixtureResponse("../sample.xml")

	songs := parseRawComment(res.TitleItems.Items[2].TID, res.TitleItems.Items[2].Comment)
	Expect(t, len(songs)).ToBe(6)

	Expect(t, songs[1].AnimeID).ToBe(SHIROBAKO)
	Expect(t, songs[1].Type).ToBe("オープニングテーマ")
	Expect(t, songs[1].Number).ToBe("2")
	Expect(t, songs[1].Title).ToBe("COLORFUL BOX")
}

func TestParseRawSubTitles(t *testing.T) {

	res := fixtureResponse("../sample.xml")

	programs := parseRawSubTitles(res.TitleItems.Items[2].TID, res.TitleItems.Items[2].SubTitles)
	Expect(t, len(programs)).ToBe(23)

	Expect(t, programs[0].AnimeID).ToBe(SHIROBAKO)
	Expect(t, programs[0].Chapter).ToBe("01")
	Expect(t, programs[0].Title).ToBe("明日に向かって、えくそだすっ！")
}

func fixtureResponse(testxml string) syobocal.TitleLookupResponse {
	f, err := os.Open(testxml)
	if err != nil {
		log.Fatalln(err)
	}
	tlr := syobocal.TitleLookupResponse{}
	decoder := xml.NewDecoder(f)
	if err := decoder.Decode(&tlr); err != nil {
		log.Fatalln(err)
	}
	return tlr
}
