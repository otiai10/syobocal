package infra

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type Db struct {
    mydb *sql.DB;
}

func GetDB(dsn string) *Db {
    prefix := "/animapi_"
    _db, err := sql.Open(
        "mysql",
        "otiai10:hoge@" + prefix + dsn,
    )
    if err != nil {
        panic(err)
    }

    defer _db.Close()

    return &Db{
        mydb: _db,
    }
}
