package anime

import "database/sql"
import "animapi/infrastructure/syobocal"

type AnimeFactory struct{}

func GetAnimeFactory() *AnimeFactory {
	return &AnimeFactory{}
}
func (f *AnimeFactory) FromRecord(record *sql.Row) *Anime {
	var id, title string
	_ = record.Scan(&id, &title)
	return &Anime{
		"0",
		title,
		"nil",
	}
}
func (f *AnimeFactory) FromSyobocalResponse(res syobocal.Response) []*Anime {
	var animes []*Anime
	for _, item := range res.TitleItems.TitleItem {
		animes = append(animes, f.item2anime(item))
	}
	return animes
}
func (f *AnimeFactory) item2anime(item syobocal.TitleItem) *Anime {
	return &Anime{
		TID:     item.TID,
		Title:   item.Title,
		Comment: item.Comment,
	}
}
