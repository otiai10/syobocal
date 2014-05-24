package application

import "fmt"

type Syobocal struct {
	name string
}

func (this Syobocal) Greet() string {
	return fmt.Sprintf(fmt_greet, this.name)
}
