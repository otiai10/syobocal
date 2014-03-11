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

type Db struct {
    mydb *sql.DB;
    shard string;
}

func GetDB(dsn string, skey string) *Db {
    prefix := "/animapi_"
    _db, err := sql.Open(
        "mysql",
        "root:@" + prefix + dsn,
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
func (db *Db)FindAllAnime() *sql.Rows {
    _rows, err := db.mydb.Query("SELECT * FROM anime_" + db.shard)
    if err != nil {
        panic(err)
    }
    return _rows
}
