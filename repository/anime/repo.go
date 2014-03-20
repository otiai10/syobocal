package repo

import . "animapi/repository"
import "animapi/model/anime"
import "animapi/infrastructure/db"
import "animapi/factory/anime"

type AnimeRepo struct {
	client RepoClient
	dsn    string
}

func NewAnimeRepo() *AnimeRepo {
	db := infra.GetDB("test", "000")
	client := RepoClient{
		Db: db,
	}
	return &AnimeRepo{
		client: client,
		dsn:    "anime",
	}
}
func AnimeRepoOf(client RepoClient) *AnimeRepo {
	return &AnimeRepo{
		client: client,
		dsn:    "anime",
	}
}

func (animeRepo *AnimeRepo) FindById(id string) *model.Anime {
	record := animeRepo.client.FindOne(
		animeRepo.dsn,
		id,
	)
	animeFacotry := factory.GetAnimeFactory()
	return animeFacotry.FromRecord(record)
}
