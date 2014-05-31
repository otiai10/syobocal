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
  tid int(11) PRIMARY KEY,
  timestamp bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8`
	_, e = table.db.Exec(query)
	return
}
func (table *ProgramsTable) Add(tid int, lastupdate int64) (e error) {
	query := `INSERT IGNORE INTO programs (tid, timestamp) VALUES (?, ?)`
	_, e = table.db.Exec(query, tid, lastupdate)
	return
}
func (table *ProgramsTable) FindSince(snc time.Duration) (rows *sql.Rows, e error) {
	timestamp := time.Now().Add(snc).Unix()
	query := fmt.Sprintf("SELECT * FROM %s WHERE timestamp > ?", table.name)
	return table.db.Query(query, timestamp)
}
func (table *ProgramsTable) Drop() (e error) {
	query := `DROP TABLE IF EXISTS programs`
	_, e = table.db.Exec(query)
	return
}
