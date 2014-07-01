package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

import "fmt"

func (m *MySQL) FindAnimes(since time.Duration) (animes []model.Anime) {
	table := infrastructure.NewAnimesTable(m.db)
	if e := table.CreateIfNotExists(); e != nil {
    fmt.Println("0001", e)
		return
	}
	rows, e := table.FindSince(since)
	if e != nil {
    fmt.Println("0002", e)
		return
	}
	return model.CreateAnimesFromMySQLResponse(rows)
}
func (m *MySQL) FindAnimesWithAnisongs(since time.Duration) (animes []model.Anime) {
	for _, anime := range m.FindAnimes(since) {
		anisongs := m.FindAnisongsByTID(anime.TID)
		anime.Anisongs = anisongs
		animes = append(animes, anime)
	}
	return animes
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
