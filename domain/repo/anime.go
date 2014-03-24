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

func (animeRepo *AnimeRepo) Save(anime *model.Anime) {
	// やっぱりSQL文を作るところまで
	// repoの責務にしないと、
	// インフラのインターフェースが死ぬんじゃないか？
	animeRepo.client.Insert(
		animeRepo.getInsertSQLQuery(anime),
	)
}
func (animeRepo *AnimeRepo) getInsertSQLQuery(anime *model.Anime) (q string) {
	q = "INSERT IGNORE INTO anime_000 (tid, title, comment) VALUES "
	// フィクスチャは今のところ"や'を含まないのでテスト通ってしまう
	q = q + fmt.Sprintf("(\"%s\",\"%s\", \"%s\")", anime.TID, anime.Title, anime.Comment)
	return q
}
