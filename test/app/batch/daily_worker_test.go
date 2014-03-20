package batch

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/app/batch"

func TestCrawlTodayAnime(t *testing.T) {
	Describe(t, "batch", func() {
		It("should get and save animes", func() {
			// tear downする

			// CrawlTodayAnimeする

			// repo.Findする

			// ある

			Expect(batch.CrawlTodayAnime()).To(Equal, true)
		})
	})
}
