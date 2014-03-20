package infra

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/infrastructure/syobocal"

// インフラの挙動を定義するテストなので、
// このテストはレポとして振る舞う
func TestGetApi(t *testing.T) {
	Describe(t, "Syobocal API", func() {
		It("can find animes カッコカリ", func() {
			api := infra.GetSyobocalAPI()
            // response := api.FindHoge()
            _ = api.FindHoge()
			Expect(true).To(Equal, true)
		})
	})
}
