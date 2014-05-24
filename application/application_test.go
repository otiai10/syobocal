package application_test

import "github.com/otiai10/animapi/application"
import "testing"
import "reflect"
import "fmt"
import "os"

func assert(t *testing.T, actual interface{}, expected interface{}) {
	if actual == expected {
		return
	}
	fmt.Printf("Expected to be `%+v`, but actual `%+v`\n", expected, actual)
	os.Exit(1)
}
func TestApplication_Get(t *testing.T) {
	syobocal := application.Get(application.SRC_SYOBOCAL)
	assert(t, reflect.TypeOf(syobocal).String(), "application.Syobocal")

	local := application.Get(application.SRC_LOCAL)
	assert(t, reflect.TypeOf(local).String(), "application.Local")
}
