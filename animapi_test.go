package animapi_test

import "github.com/otiai10/animapi"
import "testing"
import "fmt"
import "os"

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		return
	}
	fmt.Printf("Expected to be `%+v`, but actual `%+v`\n", expected, actual)
	os.Exit(1)
}
func TestAnimapi_LOCAL(t *testing.T) {
	assert(
		t,
		animapi.LOCAL.Greet(),
		"Hi, I'm Local!",
	)
}
func TestAnimapi_SYOBOCAL(t *testing.T) {
	assert(
		t,
		animapi.SYOBOCAL.Greet(),
		"Hi, I'm Syobocal!",
	)
}
