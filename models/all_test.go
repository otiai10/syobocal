package models

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/otiai10/animapi/config"
	. "github.com/otiai10/mint"
)

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
	db, err := gorm.Open("mysql", config.Values.MySQL()+"/"+config.Values.DBName())
	if err != nil {
		log.Fatalln("up", err)
	}
	db.LogMode(true)
	db.DropTable(&Anime{})

	Init(db)
}

func down() {
}

func TestInit(t *testing.T) {
	Expect(t, true).ToBe(true)

	anime := &Anime{
		ID:    1000,
		Title: "SHIROBAKO",
	}
	db := DB()
	err := db.Create(anime).Error
	Expect(t, err).ToBe(nil)
}
