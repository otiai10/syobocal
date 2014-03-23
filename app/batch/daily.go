package batch

import "animapi/infrastructure/syobocal"
import "animapi/domain/factory/anime"

import "time"
import "fmt"

func CrawlTodayAnime() bool {
	// syobocalからアニメを取得する
	syobocal := syobocal.GetAPI()

	now := time.Now()
	dur, _ := time.ParseDuration("-24h")
	yesterday := time.Now().Add(dur)

	res := syobocal.RequestByRange(yesterday, now)

	animeFactory := factory.GetAnimeFactory()
	animes := animeFactory.FromSyobocalResponse(res)

	for i, anime := range animes {
		fmt.Println(
			i,
			anime.Title,
			anime.TID,
		)
	}

	return true
}
