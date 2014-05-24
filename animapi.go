package animapi

import "time"

func Since(snc string) (dur time.Duration, e error) {
	s := &since{val: snc}
	return s.Parse()
}
