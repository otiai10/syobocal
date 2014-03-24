package batch

import "animapi/infrastructure/syobocal"
import "animapi/domain/factory/anime"

import "animapi/domain/repo"

import "time"

func CrawlTodayAnime() bool {
	// syobocalからアニメを取得する
	syobocal := syobocal.GetAPI()

	now := time.Now()
	dur, _ := time.ParseDuration("-24h")
	yesterday := time.Now().Add(dur)

	res := syobocal.RequestByRange(yesterday, now)

	animeFactory := factory.GetAnimeFactory()
	animes := animeFactory.FromSyobocalResponse(res)

	animeRepo := repo.NewAnimeRepo()
	for _, anime := range animes {
		animeRepo.Save(anime)
	}

	return true
}
