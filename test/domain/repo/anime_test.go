package repo

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/domain/repo"

func TestFindById(t *testing.T) {
	Describe(t, "FindById", func() {
		It("should find anime by id", func() {
			animeRepo := repo.NewAnimeRepo()
			anime := animeRepo.FindById("1")
			Expect(anime.Title).To(Equal, "凪のあすから")
		})
	})
}

func TestFindFromSyobocal(t *testing.T) {
	Describe(t, "FindFromSyobocal", func() {
		It("should find anime from syobocal with query", func() {
			dummyClient := repo.DummyHTTPClient{}
			animeRepo := repo.AnimeRepoOf(dummyClient)
			animeRepo.FindFromSyobocal("20140320_000000", "")
		})
	})
}
