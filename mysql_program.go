package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

func (m *MySQL) FindPrograms(since time.Duration) (programs []model.Program) {
	table := infrastructure.NewProgramsTable(m.db)
	rows, e := table.FindSince(since)
	if e != nil {
		return
	}
	return model.CreateProgramsFromMySQLResponse(rows)
}
