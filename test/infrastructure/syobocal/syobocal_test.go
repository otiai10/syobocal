package infra

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/infrastructure/syobocal"
import (
	"io/ioutil" // fixtureモジュールが必要かも
)

// インフラの挙動を定義するテストなので、
// このテストはレポとして振る舞う
func TestGetApi(t *testing.T) {
	Describe(t, "Syobocal API", func() {
		It("can find animes カッコカリ", func() {
			// api := infra.GetSyobocalAPI()
			api := infra.SyobocalApiOf(
				DummyHTTPClient{
					baseURL: "this.is.dummy.client",
				},
			)
			// response := api.FindHoge()
			_ = api.FindHoge()
			Expect(true).To(Equal, true)
		})
	})
}

type DummyHTTPClient struct {
	*infra.SyobocalHTTPClient
	baseURL string
}

func (client DummyHTTPClient) FindHoge() []byte {
	var xml []byte
	xml, e := ioutil.ReadFile("../../fixture/syobocal/syobocal.response.xml")
	if e != nil {
		panic(e)
	}
	return xml
}
