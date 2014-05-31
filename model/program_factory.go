package model

import "database/sql"

func CreateProgramsFromMySQLResponse(rows *sql.Rows) (programs []Program) {
	for rows.Next() {
		var tid int
		var timestamp int64
		rows.Scan(&tid, &timestamp)
		program := Program{
			Anime{
				TID:         tid,
				LastUpdated: timestamp,
			},
		}
		programs = append(programs, program)
	}
	return
}
