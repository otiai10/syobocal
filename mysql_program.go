package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

func (m *MySQL) FindPrograms(since time.Duration) (programs []model.Program) {
	table := infrastructure.NewProgramsTable(m.db)
	if e := table.CreateIfNotExists(); e != nil {
		return
	}
	rows, e := table.FindSince(since)
	if e != nil {
		return
	}
	return model.CreateProgramsFromMySQLResponse(rows)
}
func (m *MySQL) AddPrograms(programs []model.Program) (e error) {
	table := infrastructure.NewProgramsTable(m.db)
	if e = table.CreateIfNotExists(); e != nil {
		return
	}
	for _, program := range programs {
		if e = table.Add(program.Anime.TID, time.Now().Unix()); e != nil {
			return
		}
	}
	return
}
func (m *MySQL) DeleteProgram(program model.Program) (e error) {
	table := infrastructure.NewProgramsTable(m.db)
	if e = table.CreateIfNotExists(); e != nil {
		return
	}
	return table.Delete(program.Anime.TID, program.Timestamp)
}
