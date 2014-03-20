package repo

import . "animapi/domain/repository"
import "animapi/domain/model/anime"
import "animapi/domain/infrastructure/db"
import "animapi/domain/factory/anime"

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
