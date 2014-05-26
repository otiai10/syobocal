package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

type syobocal struct{}

var SYOBOCAL = syobocal{}

func (s syobocal) Greet() string {
	return "Hi, I'm Syobocal!"
}
func (s syobocal) FindPrograms(since time.Duration) []model.Program {
	client := infrastructure.NewSyobocalClient()
	res, _ := client.TitleLookup(since)
	return model.CreateProgramsFromSyobocalResponse(res)
}
