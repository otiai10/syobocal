package infrastructure

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "time"
import "fmt"

// アクティブテーブルへの参照を返す
func NewProgramsTable(db *sql.DB) *ProgramsTable {
	return &ProgramsTable{
		db:   db,
		name: "programs",
	}
}

type ProgramsTable struct {
	db   *sql.DB
	name string
}

func (table *ProgramsTable) CreateIfNotExists() (e error) {
	query := `
CREATE TABLE IF NOT EXISTS programs (
  tid int(11) DEFAULT NULL,
  timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8`
	_, e = table.db.Exec(query)
	return
}
func (table *ProgramsTable) FindSince(snc time.Duration) (rows *sql.Rows, e error) {
	timestamp := time.Now().Add(snc).Unix()
	query := fmt.Sprintf("SELECT * FROM %s WHERE timestamp < ?", table.name)
	return table.db.Query(query, timestamp)
}
