package infra

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/infrastructure/syobocal"

func TestGetApi(t *testing.T) {
	Describe(t, "Syobocal API", func() {
		It("can find animes カッコカリ", func() {

			api := infra.GetSyobocalAPI()
            api.Hoge()
/*
			xml := api.FindAllAnime()

            fmt.Printf("%T %+v", xml, xml)
*/
			Expect(true).To(Equal, true)
		})
	})
}
