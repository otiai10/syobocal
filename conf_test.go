package animapi_test

import "github.com/otiai10/animapi"
import "testing"

func TestConf(t *testing.T) {
	conf := animapi.File("my.conf")
	assert(t, conf.Err, nil)
	assert(t, conf.Port, "3306")

	conf = animapi.File("my.conf", "test")
	assert(t, conf.Port, "2222")
}
