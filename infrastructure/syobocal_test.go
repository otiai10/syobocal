package infrastructure_test

import "github.com/otiai10/animapi/infrastructure"
import "testing"
import "time"
import . "github.com/otiai10/mint"

func TestSyobocalHttpClient(t *testing.T) {
	since, _ := time.ParseDuration("-24h")
	client := infrastructure.NewSyobocalClient()
	res, _ := client.TitleLookup(since)
	for _, item := range res.TitleItems.Items {
		Expect(t, item).TypeOf("infrastructure.TitleItem")
	}
}
