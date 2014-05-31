package model

import "database/sql"

func CreateAnisongsFromMySQLResponse(rows *sql.Rows) (anisongs []Anisong) {
	for rows.Next() {
		var tid int
		var title string
		var label string
		var indx int
		var detail string
		rows.Scan(&tid, &title, &label, &indx, &detail)
		anisong := CreateAnisong(tid, title, label, indx, detail)
		anisongs = append(anisongs, anisong)
	}
	return
}
func CreateAnisong(tid int, title string, label string, indx int, detail string) Anisong {
	return Anisong{
		TID:    tid,
		Title:  title,
		Label:  label,
		Index:  indx,
		Detail: detail,
	}
}
