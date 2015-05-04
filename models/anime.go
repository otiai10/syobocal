package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// Anime アニメです.
type Anime struct {
	ID             int            `db:"id" syobocal:"TID"`
	UpdatedAt      time.Time      `db:"updated_at" syobocal:"LastUpdated"`
	Title          string         `db:"title" syobocal:"Title"`
	CommentRaw     string         `db:"comment" syobocal:"Comment"`
	Category       Category       `db:"category" syobocal:"Category"`
	FirstBroadcast time.Time      `db:"first_broadcast" syobocal:"FirtYear_FirstMonth"`
	FirstEnded     mysql.NullTime `db:"first_ended" syobocal:"FirstEndYear_FirstEndMonth"`
	Keywords       []string       `db:"-" syobocal:"Keywords,_comma_splitted"`
	Songs          []Song         `db:"-"`
	Programs       []Program      `db:"-"`
}

// Save ...
func (anime *Anime) Save(db *gorp.DbMap) error {
	return db.Insert(anime)
}
