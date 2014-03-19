package repo

import . "animapi/repository"
import "animapi/model/anime"
import "animapi/infrastructure/db"

import "fmt"

type AnimeRepo struct {
    client RepoClient
    dsn string
}

func NewAnimeRepo() *AnimeRepo {
    db := infra.GetDB("test", "000")
    client := RepoClient{
        Db: db,
    }
    return &AnimeRepo{
        client: client,
        dsn: "anime",
    }
}
func AnimeRepoOf(client RepoClient) *AnimeRepo {
    return &AnimeRepo{
        client: client,
        dsn: "anime",
    }
}

func (animeRepo *AnimeRepo)FindById(id string) *model.Anime {
    // とりあえずハード
    row := animeRepo.client.FindOne(
        animeRepo.dsn,
        id,
    )
    // {{{ TODO: Factoryを呼ぶ
    var title string
    _ = row.Scan(&id, &title)
    fmt.Println(
        id,
        title,
    )
    // }}}
    return &model.Anime{
        title,
    }
}
