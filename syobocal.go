package animapi

type syobocal struct{}

var SYOBOCAL = syobocal{}

func (s syobocal) Greet() string {
	return "Hi, I'm Syobocal!"
}
