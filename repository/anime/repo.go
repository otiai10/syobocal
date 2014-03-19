package repo

import . "animapi/repository"
import "animapi/model/anime"

type AnimeRepo struct {
    client RepoClient
}

func NewAnimeRepo() *AnimeRepo {
    client := RepoClient{}
    return &AnimeRepo{
        client: client,
    }
}
func AnimeRepoOfTest() *AnimeRepo {
    client := RepoClient{}
    return &AnimeRepo{
        client: client,
    }
}

func (animeRepo *AnimeRepo)FindById(id string) *model.Anime {
    // とりあえずハード
    return &model.Anime{
        "いなり、こんこん、恋いろは。",
    }
}
