/**
 * データベースとの接続を担保する
 * ここはインフラ層であり
 * アプリケーションのことは
 * いっさい、これっぽっちも
 * 関知してはいけない
 * 逆に言えば、データベースが
 * mongoDBに移行することになっても
 * 動くように意識してつくる
 * そのへんを吸収する
 */
package infra

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

type Db struct {
	mydb  *sql.DB
	shard string
}

func GetDB(dsn string, skey string) *Db {
	prefix := "/animapi_"
	_db, err := sql.Open(
		"mysql",
		"root:@"+prefix+dsn,
	)
	if err != nil {
		panic(err)
	}

	// defer _db.Close()

	return &Db{
		mydb:  _db,
		shard: skey,
	}
}

// とりあえずテスト
// 本当はここで「アニメ」という言葉を使えない
func (db *Db) FindAllAnime() *sql.Rows {
	query := "SELECT id, title FROM anime_" + db.shard
	fmt.Println("[find]\t" + query)
	_rows, err := db.mydb.Query(query)
	if err != nil {
		panic(err)
	}
	return _rows
}
func (db *Db) InsertAnime() sql.Result {
	statement := `INSERT INTO anime_%s
                  (tid, title, firstYear, firstMonth, firstEndYear, firstEndMonth) VALUE
                  (0, '凪のあすから', 2013, 12, 2014, 3)`
	q := fmt.Sprintf(statement, db.shard)
	return db.exec(q)
}
func (db *Db) DropAnimeTable() sql.Result {
	q := "DROP TABLE anime_" + db.shard
	return db.exec(q)
}
func (db *Db) CreateAnimeTable() sql.Result {
	statement := `CREATE TABLE IF NOT EXISTS anime_%s (
                    id INTEGER UNIQUE AUTO_INCREMENT,
                    tid INT UNSIGNED NOT NULL DEFAULT 0,
                    title TEXT NOT NULL,
                    comment TEXT,
                    firstYear INTEGER NOT NULL,
                    firstMonth INTEGER NOT NULL,
                    firstEndYear INTEGER NOT NULL,
                    firstEndMonth INTEGER NOT NULL
                  )`
	q := fmt.Sprintf(statement, db.shard)
	return db.exec(q)
}
func (db *Db) exec(query string) sql.Result {
	_res, err := db.mydb.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("[exec]\t" + query)
	return _res
}
func (db *Db) queryOne(query string) *sql.Row {
	return db.mydb.QueryRow(query)
}

// こっから下はマジ
func (db *Db) FindOne(tableName string, id string) *sql.Row {
	query := "SELECT id,title FROM " + tableName + "_" + db.shard + " WHERE id='" + id + "'"
	_row := db.queryOne(query)
	return _row
}

func (db *Db) Insert(sql string) {
	res, e := db.mydb.Exec(sql)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%+v", res)
}

func (db *Db) Exec(query string, args ...interface{}) {
	res, e := db.mydb.Exec(query, args...)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%+v", res)
}
