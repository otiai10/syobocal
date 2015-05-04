package models

// Song アニソン的な.
type Song struct {
	ID         int
	AnimeID    int    `db:"index"`
	Type       string `db:"song_type"` // 基本的には"オープニング","エンディング"
	Number     string `db:"seq"`
	Title      string
	Attributes map[string]string
}
