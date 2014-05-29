package animapi_test

import "github.com/otiai10/animapi"
import "testing"
import . "github.com/otiai10/mint"

func TestAnimapi_SYOBOCAL(t *testing.T) {
	Expect(t, animapi.SYOBOCAL.Greet()).ToBe("Hi, I'm Syobocal!")
}
func TestAnimapi_SYOBOCAL_FindPrograms(t *testing.T) {
	since, _ := animapi.Since("-4h")
	programs := animapi.SYOBOCAL.FindPrograms(since)
	Expect(t, programs).TypeOf("[]model.Program")
}
func TestAnimapi_SYOBOCAL_FindProgramsSince(t *testing.T) {
	programs, e := animapi.SYOBOCAL.FindProgramsSince("-1w")
	Expect(t, e).ToBe(nil)
	Expect(t, programs).TypeOf("[]model.Program")
}
