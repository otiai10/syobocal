package repo

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/domain/repository/anime"

func TestFindById(t *testing.T) {
	Describe(t, "FindById", func() {
		It("should find anime by id", func() {
			animeRepo := repo.NewAnimeRepo()
			anime := animeRepo.FindById("1")
			Expect(anime.Title).To(Equal, "凪のあすから")
		})
	})
}
