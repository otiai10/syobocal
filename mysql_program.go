package animapi

import "time"
import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

func (m *MySQL) FindPrograms(since time.Duration) []model.Program {
	table := infrastructure.NewProgramTable(m.db)
	res, _ := table.Find(since)
	return model.CreateProgramsFromMySQLResponse(res)
}
