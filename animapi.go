package animapi

import "time"

// TODO: Sinceっていう名前もイケてないし、
// そもそもsncをParseDurationできるんだったらそれでよくね？
func Since(snc string) (dur time.Duration, e error) {
	s := &since{val: snc}
	return s.Parse()
}
