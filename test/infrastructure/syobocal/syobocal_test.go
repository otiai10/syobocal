package infra

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
)

// Test this
import "animapi/infrastructure/syobocal"

// 良い感じ 2014/03/22
func TestRequestByRange(t *testing.T) {
	Describe(t, "RequestByRange", func() {
		It("should find animes by range", func() {
			syobocal := syobocal.GetAPI()
			res := syobocal.RequestByRange("20140322_214000", "")
			typeof := fmt.Sprintf("%T", res)
			Expect(typeof).To(Equal, "syobocal.Response")
		})
	})
}
