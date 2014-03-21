package infra

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
)

// Test this
import "animapi/infrastructure/syobocal"

// インフラの挙動を定義するテストなので、
// このテストはレポとして振る舞う
func TestGetApi(t *testing.T) {
	Describe(t, "RequestQuery", func() {
		It("should find animes from syobocal by query", func() {
			client := &infra.SyobocalDummyHTTPClient{}
			api := infra.SyobocalApiOf(client)
			res := api.RequestQuery("20140321", "20140324")
			// reflect使うまでもない
			typeof := fmt.Sprintf("%T", res)
			Expect(typeof).To(Equal, "model.SyobocalResponse")
		})
	})
}
