package animapi_test

import "github.com/otiai10/animapi"
import "testing"
import . "github.com/otiai10/mint"

func TestAnimapi_DB(t *testing.T) {
	mysqlClient := animapi.DB("./my.conf")
	Expect(t, mysqlClient.Err).ToBe(nil)

	mysqlClient = animapi.DB("./notfound.conf")
	Expect(t, mysqlClient.Err.Error()).ToBe("open ./notfound.conf: no such file or directory")

	mysqlClient = animapi.DB("./my.conf", "test")
	Expect(t, mysqlClient.Err).ToBe(nil)
}

func TestAnimapi_DB_FindPrograms(t *testing.T) {
	since, _ := animapi.Since("-4h")
	c := "./my.conf"
	programs := animapi.DB(c, "test").FindPrograms(since)
	Expect(t, programs).TypeOf("[]model.Program")
}
