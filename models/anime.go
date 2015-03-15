package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// Category しょぼかるといっしょ.
type Category int

// Anime アニメです.
type Anime struct {
	ID             int            `sql:"id,not null;unique" syobocal:"TID"`
	UpdatedAt      time.Time      `sql:"updated_at" syobocal:"LastUpdated"`
	Title          string         `sql:"title,not null;unique" syobocal:"Title"`
	CommentRaw     string         `syobocal:"Comment"`
	Category       Category       `syobocal:"Category"`
	FirstBroadcast time.Time      `syobocal:"FirtYear_FirstMonth"`
	FirstEnded     mysql.NullTime `sql:"default null" syobocal:"FirstEndYear_FirstEndMonth"`
	Keywords       []string       `syobocal:"Keywords,_comma_splitted"`
	Songs          []Song
	Programs       []Program
}

// Song アニソン的な.
type Song struct {
	ID         int
	AnimeID    int    `sql:"index"`
	Type       string `sql:"song_type"` // 基本的には"オープニング","エンディング"
	Number     string `sql:"seq"`
	Title      string
	Attributes map[string]string
}

// Program しょぼかるのProgに相当.
type Program struct {
	ID      int
	AnimeID int `sql:"index"`
	Title   string
	Chapter string
}
