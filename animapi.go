package animapi

import "github.com/otiai10/animapi/application"
import "time"

var (
	LOCAL    = application.Get(application.SRC_LOCAL)
	SYOBOCAL = application.Get(application.SRC_SYOBOCAL)
)

func Since(snc string) (dur time.Duration, e error) {
	s := &since{val: snc}
	return s.Parse()
}
