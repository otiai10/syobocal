package infrastructure_test

import "github.com/otiai10/animapi/infrastructure"
import "testing"
import "fmt"
import "os"
import "time"
import "reflect"

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		return
	}
	fmt.Printf("Expected to be `%+v`, but actual `%+v`\n", expected, actual)
	os.Exit(1)
}

func TestSyobocalHttpClient(t *testing.T) {
	since, _ := time.ParseDuration("-24h")
	client := infrastructure.NewSyobocalClient()
	res, _ := client.TitleLookup(since)
	for _, item := range res.TitleItems.Items {
		assert(t, reflect.TypeOf(item).String(), "infrastructure.TitleItem")
	}
}
