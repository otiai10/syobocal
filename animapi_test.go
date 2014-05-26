package animapi_test

import "github.com/otiai10/animapi"
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
func TestAnimapi_SYOBOCAL(t *testing.T) {
	assert(
		t,
		animapi.SYOBOCAL.Greet(),
		"Hi, I'm Syobocal!",
	)
}
func TestAnimapi_SYOBOCAL_FindPrograms(t *testing.T) {
	since, _ := animapi.Since("-4h")
	programs := animapi.SYOBOCAL.FindPrograms(since)
	assert(
		t,
		reflect.TypeOf(programs).String(),
		"[]model.Program",
	)
}
func TestAnimapi_SYOBOCAL_FindProgramsSince(t *testing.T) {
	programs, e := animapi.SYOBOCAL.FindProgramsSince("-1w")
	assert(t, e, nil)
	assert(
		t,
		reflect.TypeOf(programs).String(),
		"[]model.Program",
	)
}
func TestAnimapi_DB(t *testing.T) {

	mysqlClient := animapi.DB("./my.conf")
	assert(t, mysqlClient.Err, nil)
	mysqlClient = animapi.DB("./notfound.conf")
	assert(
		t,
		mysqlClient.Err.Error(),
		"open ./notfound.conf: no such file or directory",
	)
	mysqlClient = animapi.DB("./my.conf", "test")
	assert(t, mysqlClient.Err, nil)
}
