package repo

import "animapi/domain/model/anime"
import "animapi/infrastructure/db"
import "animapi/domain/factory/anime"

import "fmt"

type AnimeRepo struct {
	client IRepoClient
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
func AnimeRepoOf(client IRepoClient) *AnimeRepo {
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

func (animeRepo *AnimeRepo) FindFromSyobocal(from string, to string) /* []*model.Anime */ {
	query := &SyobocalQuery{
		"TitleLookup",
		from,
		to,
	}
	var xml []byte
	xml = animeRepo.client.Query(query)
	fmt.Println(xml)
}
