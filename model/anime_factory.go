package model

import "database/sql"

func CreateAnimesFromMySQLResponse(rows *sql.Rows) (animes []Anime) {
	for rows.Next() {
		var tid int
		var title string
		var lastUpdated int64
		rows.Scan(&tid, &title, &lastUpdated)
		animes = append(animes, CreateAnime(tid, title, lastUpdated))
	}
	return
}
func CreateAnime(tid int, title string, lastUpdated int64) Anime {
	return Anime{
		TID:         tid,
		Title:       title,
		LastUpdated: lastUpdated,
	}
}
