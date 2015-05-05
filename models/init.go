package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var _db *gorp.DbMap

const (
	createdatabase = "CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8"
)

// Init データベースを初期化するし、gorp的にテーブルつくる
func Init(mysqlURI string, databasename string) *gorp.DbMap {

	db, err := sql.Open("mysql", mysqlURI+"/")
	onError(err, "initial open to /")

	dbmap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "utf8",
		},
	}

	_, err = dbmap.Exec(fmt.Sprintf(createdatabase, databasename))
	onError(err, "database create")
	db, err = sql.Open("mysql", mysqlURI+"/"+databasename)
	onError(err, "open database")
	dbmap.Db = db

	dbmap.AddTableWithName(Anime{}, "animes")
	// .SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	onError(err, "create table")

	return dbmap
}

func onError(err error, msg string) {
	if err == nil {
		return
	}
	log.Fatalln(msg, err)
}
