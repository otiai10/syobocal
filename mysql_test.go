package animapi_test

import "github.com/otiai10/animapi"
import "testing"
import . "github.com/otiai10/mint"

func TestMySQL(t *testing.T) {
	client := animapi.DB("./my.conf", "test")
	Expect(t, client).TypeOf("*animapi.MySQL")
}
func TestMySQL_FindProgramsSince(t *testing.T) {
	programs, e := animapi.DB("./my.conf", "test").FindProgramsSince("-1w")
	Expect(t, programs).TypeOf("[]model.Program")
	Expect(t, e).ToBe(nil)
}
