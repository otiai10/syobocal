package infra

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

import "animapi/domain/infra/db"

func TestGetDb(t *testing.T) {
	Describe(t, "Db", func() {
		It("can read sql カッコカリ", func() {

			db := infra.GetDB("test", "000")

			_ = db.CreateAnimeTable()
			_ = db.DropAnimeTable()
			_ = db.CreateAnimeTable()

			_ = db.InsertAnime()

			rows := db.FindAllAnime()
			rows.Next()
			var title string
			var id string
			e := rows.Scan(&id, &title)
			if e != nil {
				panic(e)
			}
			Expect(id).To(Equal, "1")
		})
	})
}
