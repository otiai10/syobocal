package application

import "fmt"

type Local struct {
	name string
}

func (this Local) Greet() string {
	return fmt.Sprintf(fmt_greet, this.name)
}
