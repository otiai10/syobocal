package animapi

import "time"

type syobocal struct{}

var SYOBOCAL = syobocal{}

func (s syobocal) Greet() string {
	return "Hi, I'm Syobocal!"
}
func (s syobocal) FindPrograms(since time.Duration) {

}
