package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

type proxySyobocal struct{}

var SYOBOCAL = proxySyobocal{}

func (s proxySyobocal) Greet() string {
	return "Hi, I'm Syobocal!"
}
func (s proxySyobocal) FindPrograms(since time.Duration) []model.Program {
	client := infrastructure.NewSyobocalClient()
	res, _ := client.TitleLookup(since)
	return model.CreateProgramsFromSyobocalResponse(res)
}
func (s proxySyobocal) FindProgramsSince(snc string) (programs []model.Program, e error) {
	var dur time.Duration
	if dur, e = Since(snc); e != nil {
		return programs, e
	}
	programs = s.FindPrograms(dur)
	return
}
