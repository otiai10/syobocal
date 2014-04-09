package infra

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
	"time"
)

// Test this
import "animapi/infrastructure/syobocal"

// 良い感じ 2014/03/22
func TestRequestByRange(t *testing.T) {
	Describe(t, "RequestByRange", func() {
		It("should find animes by range", func() {
			// syobocal := syobocal.GetAPI()
			syobocal := syobocal.GetAPIofTest()
			dur, _ := time.ParseDuration("-3h")
			from := time.Now().Add(dur)
			to := time.Now()
			res := syobocal.RequestByRange(from, to)
			typeof := fmt.Sprintf("%T", res)
			Expect(typeof).To(Equal, "syobocal.Response")
		})
	})
}
