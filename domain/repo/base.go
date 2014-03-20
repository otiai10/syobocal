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
import "animapi/domain/infra/db"
import "database/sql"

var Repo interface{}

type RepoClient struct {
	Db *infra.Db
}

func (client RepoClient) FindOne(dsn, id string) *sql.Row {
	return client.Db.FindOne(dsn, id)
}
