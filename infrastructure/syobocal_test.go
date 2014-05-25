package infrastructure_test

import "github.com/otiai10/animapi/infrastructure"
import "testing"
import "fmt"
import "os"
import "time"

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		return
	}
	fmt.Printf("Expected to be `%+v`, but actual `%+v`\n", expected, actual)
	os.Exit(1)
}

func TestSyobocalHttpClient(t *testing.T) {
	since, _ := time.ParseDuration("-2h")
	client := infrastructure.NewSyobocalClient()
	res, e := client.TitleLookup(since)
	fmt.Printf("%+v\n%+v\n", res, e)
}
