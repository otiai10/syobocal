package models

// Song ...
type Song struct {
	Title    string   `json:"title"`
	Label    string   `json:"label"`    // オープニングとかエンディングとか挿入歌とか
	Words    []string `json:"words"`    // 作詞
	Music    []string `json:"music"`    // 作曲
	Composer []string `json:"composer"` // 編曲
	Singer   []string `json:"singer"`   // 歌
	Anime    string   `json:"anime"`    // アニメタイトル
}
