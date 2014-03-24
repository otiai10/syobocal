package repo

/**
 * レポジトリは
 * データへどんな処理をするか定義し
 * アプリケーションによるデータへのインターフェースを提供する
 * 各レポジトリオブジェクトのメソッドは
 * データベースがどんな形式であるか（infra.db）
 * データベースからの返りが何か（factory）
 * 全く知らないように作る
 */
import "animapi/infrastructure/db"
import "database/sql"

var Repo interface{}

type IRepoClient interface {
	FindOne(dsn, id string) *sql.Row
	Insert(sql string)
	Exec(query string, args ...interface{})
}
type RepoClient struct {
	Db *infra.Db
}

func (client RepoClient) FindOne(dsn, id string) *sql.Row {
	return client.Db.FindOne(dsn, id)
}
func (client RepoClient) Insert(sql string) {
	client.Db.Insert(sql)
}
func (client RepoClient) Exec(query string, args ...interface{}) {
	client.Db.Exec(query, args...)
}
