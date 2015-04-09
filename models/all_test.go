package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/otiai10/animapi/config"
	. "github.com/otiai10/mint"
	"gopkg.in/gorp.v1"
)

var _dbmap *gorp.DbMap

func init() {
	config.Init("test")
}

func TestMain(m *testing.M) {
	up()
	ret := m.Run()
	down()
	os.Exit(ret)
}

func up() {
	db, err := sql.Open("mysql", config.Values.MySQL()+"/")
	if err != nil {
		log.Fatalln("up1", err)
	}
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", config.Values.DBName()))
	if err != nil {
		log.Fatalln("up2", err)
	}
	_dbmap = Init(config.Values.MySQL(), config.Values.DBName())
}

func down() {
}

func TestInit(t *testing.T) {
	Expect(t, true).ToBe(true)

	anime := &Anime{
		ID:    1000,
		Title: "SHIROBAKO",
	}
	err := anime.Save(_dbmap)
	Expect(t, err).ToBe(nil)
}
