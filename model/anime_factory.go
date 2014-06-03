package model

import "database/sql"

func CreateAnimesFromMySQLResponse(rows *sql.Rows) (animes []Anime) {
	for rows.Next() {
		var tid int
		var title string
		var lastUpdated int64
		var category int
		rows.Scan(&tid, &title, &lastUpdated, &category)
		animes = append(animes, CreateAnime(tid, title, lastUpdated, category))
	}
	return
}
func CreateAnime(tid int, title string, lastUpdated int64, category int) Anime {
	return Anime{
		TID:         tid,
		Title:       title,
		LastUpdated: lastUpdated,
		Category:    category,
	}
}
