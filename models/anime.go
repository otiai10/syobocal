package models

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// Anime アニメです.
type Anime struct {
	ID             int            `json:"id" gorm:"column:id;primary_key" sql:"UNIQUE;AUTO_INCREMENT:NO;NOT NULL"` // しょぼかる的なIDを入れる
	CreatedAt      time.Time      // レコードのcreated_at
	UpdatedAt      time.Time      // レコードのupdated_at
	Title          string         // アニメタイトル
	CommentRaw     string         // コメント（生）
	Category       Category       // カテゴリー
	FirstBroadcast time.Time      // 最初に放映された時間
	FirstEnded     mysql.NullTime // 最初に放映終了した時間
	Keywords       []string       `sql:"-"` // しょぼかる的にはKeywords
	Songs          []*Song        `sql:"-"` // しょぼかる的にはCommentに入ってる
	Programs       []*Program     `sql:"-"` // しょぼかる的にはSubTitlesとよばれている
	Links          []*Link        `sql:"-"` // しょぼかる的にはCommentに入ってる
}
