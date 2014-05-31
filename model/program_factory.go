package model

import "database/sql"
import "github.com/otiai10/animapi/infrastructure"

func CreateProgramsFromMySQLResponse(rows *sql.Rows) (programs []Program) {
	for rows.Next() {
		var tid int
		var timestamp int64
		rows.Scan(&tid, &timestamp)
		program := Program{
			Anime: Anime{
				TID:         tid,
				LastUpdated: timestamp,
			},
			Timestamp: timestamp,
		}
		programs = append(programs, program)
	}
	return
}

func CreateProgramsFromSyobocalResponse(response infrastructure.SyobocalResponse) []Program {
	programs := []Program{}
	for _, item := range response.TitleItems.Items {
		if anime, e := CreateAnimeFromTitelItem(item); e == nil {
			programs = append(programs, CreateProgram(anime))
		}
	}
	return programs
}
func CreateProgram(anime Anime) Program {
	return Program{
		Anime: anime,
	}
}
