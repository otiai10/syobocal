package infrastructure

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

// アクティブテーブルへの参照を返す
func NewAnisongsTable(db *sql.DB) *AnisongsTable {
	return &AnisongsTable{
		db:   db,
		name: "anisongs",
	}
}

type AnisongsTable struct {
	db   *sql.DB
	name string
}

func (table *AnisongsTable) CreateIfNotExists() (e error) {
	query := `
CREATE TABLE IF NOT EXISTS anisongs (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  tid int(11) NOT NULL,
  title text NOT NULL,
  label text,
  indx int,
  detail text
) ENGINE=InnoDB DEFAULT CHARSET=utf8`
	_, e = table.db.Exec(query)
	return
}
func (table *AnisongsTable) Add(tid int, title string, label string, indx int, detail string) (e error) {
	query := `INSERT IGNORE INTO anisongs (tid, title, label, indx, detail) VALUES (?, ?, ?, ?, ?)`
	_, e = table.db.Exec(query, tid, title, label, indx, detail)
	return
}
func (table *AnisongsTable) FindByTID(tid int) (rows *sql.Rows, e error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE tid=?", table.name)
	return table.db.Query(query, tid)
}
func (table *AnisongsTable) DeleteByTID(tid int) (e error) {
	query := `DELETE FROM anisongs WHERE tid=?`
	_, e = table.db.Exec(query, tid)
	return
}
func (table *AnisongsTable) Drop() (e error) {
	query := `DROP TABLE IF EXISTS anisongs`
	_, e = table.db.Exec(query)
	return
}
