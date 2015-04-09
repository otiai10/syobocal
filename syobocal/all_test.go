package syobocal

import (
	"testing"
	"time"

	. "github.com/otiai10/mint"
)

func TestClient(t *testing.T) {
	Expect(t, true).ToBe(true)
	client := NewClient()
	Expect(t, client).TypeOf("*syobocal.Client")
}

func TestClient_TitleLookup(t *testing.T) {
	Expect(t, true).ToBe(true)
	client := NewClient()
	res1, err := client.TitleLookup().LastUpdate(
		time.Now().Add(-24*time.Hour), time.Now(),
	).Do()

	Expect(t, res1.Result.Code).ToBe(200)
	Expect(t, err).ToBe(nil)

	res2, err := client.TitleLookup().From(time.Now().Add(-2 * time.Hour)).Do()

	Expect(t, len(res2.TitleItems.Items) < len(res1.TitleItems.Items)).ToBe(true)
}
