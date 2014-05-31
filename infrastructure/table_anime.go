package infrastructure

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "time"
import "fmt"

// アクティブテーブルへの参照を返す
func NewAnimesTable(db *sql.DB) *AnimesTable {
	return &AnimesTable{
		db:   db,
		name: "animes",
	}
}

type AnimesTable struct {
	db   *sql.DB
	name string
}

func (table *AnimesTable) CreateIfNotExists() (e error) {
	query := `
CREATE TABLE IF NOT EXISTS animes (
  tid int(11) NOT NULL,
  title text NOT NULL,
  lastUpdated bigint NOT NULL,
  primary key(tid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8`
	_, e = table.db.Exec(query)
	return
}
func (table *AnimesTable) Add(tid int, title string, lastupdate int64) (e error) {
	query := `INSERT IGNORE INTO animes (tid, title, lastUpdated) VALUES (?, ?, ?)`
	_, e = table.db.Exec(query, tid, title, lastupdate)
	return
}
func (table *AnimesTable) FindSince(snc time.Duration) (rows *sql.Rows, e error) {
	lastUpdated := time.Now().Add(snc).Unix()
	query := fmt.Sprintf("SELECT * FROM %s WHERE lastUpdated > ?", table.name)
	return table.db.Query(query, lastUpdated)
}
func (table *AnimesTable) Delete(tid int) (e error) {
	query := `DELETE FROM animes WHERE tid=?`
	_, e = table.db.Exec(query, tid)
	return
}
func (table *AnimesTable) Drop() (e error) {
	query := `DROP TABLE IF EXISTS animes`
	_, e = table.db.Exec(query)
	return
}
