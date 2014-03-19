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
	q := "INSERT INTO anime_" + db.shard + " (title) VALUE ('凪のあすから')"
	return db.exec(q)
}
func (db *Db) DropAnimeTable() sql.Result {
	q := "DROP TABLE anime_" + db.shard
	return db.exec(q)
}
func (db *Db) CreateAnimeTable() sql.Result {
	q := "CREATE TABLE IF NOT EXISTS anime_" + db.shard + " (id INTEGER UNIQUE AUTO_INCREMENT, title TEXT NOT NULL)"
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
