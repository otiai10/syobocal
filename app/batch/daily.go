package batch

import "animapi/infrastructure/syobocal"
import "animapi/domain/factory/anime"

import "fmt"

func CrawlTodayAnime() bool {
	// syobocalからアニメを取得する
	syobocal := syobocal.GetAPI()
	res := syobocal.RequestByRange("20140322_214050", "")
	animeFactory := factory.GetAnimeFactory()
	animes := animeFactory.FromSyobocalResponse(res)

	fmt.Printf("%+v", animes)

	return true
}
