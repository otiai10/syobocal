package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

type SyobocalClient struct{}

var SYOBOCAL = SyobocalClient{}

func (s SyobocalClient) Greet() string {
	return "Hi, I'm Syobocal!"
}
func (s SyobocalClient) FindPrograms(since time.Duration) []model.Program {
	client := infrastructure.NewSyobocalClient()
	res, _ := client.TitleLookup(since)
	return model.CreateProgramsFromSyobocalResponse(res)
}
func (s SyobocalClient) FindProgramsSince(snc string) (programs []model.Program, e error) {
	var dur time.Duration
	if dur, e = Since(snc); e != nil {
		return programs, e
	}
	programs = s.FindPrograms(dur)
	return
}
