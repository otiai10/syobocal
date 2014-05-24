package animapi

import "github.com/otiai10/animapi/application"

var (
	LOCAL    = application.Get(application.SRC_LOCAL)
	SYOBOCAL = application.Get(application.SRC_SYOBOCAL)
)
