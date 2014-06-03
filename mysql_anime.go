package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

func (m *MySQL) FindAnimes(since time.Duration) (animes []model.Anime) {
	table := infrastructure.NewAnimesTable(m.db)
	if e := table.CreateIfNotExists(); e != nil {
		return
	}
	rows, e := table.FindSince(since)
	if e != nil {
		return
	}
	return model.CreateAnimesFromMySQLResponse(rows)
}
func (m *MySQL) AddAnime(anime model.Anime) (e error) {
	table := infrastructure.NewAnimesTable(m.db)
	if e = table.CreateIfNotExists(); e != nil {
		return
	}
	return table.Add(anime.TID, anime.Title, anime.LastUpdated, anime.Category)
}
func (m *MySQL) DeleteAnime(anime model.Anime) (e error) {
	table := infrastructure.NewAnimesTable(m.db)
	if e = table.CreateIfNotExists(); e != nil {
		return
	}
	return table.Delete(anime.TID)
}
