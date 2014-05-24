package animapi_test

import "github.com/otiai10/animapi"
import "testing"
import "fmt"

func TestConf(t *testing.T) {
	conf := animapi.File("conf.my")
	fmt.Printf("%+v", conf)
}
