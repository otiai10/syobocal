package anime

import "time"

// Anime DBにEncode/Decodeしたいので、なるべくプリミティブがいい。
type Anime struct {
	SID         int       `json:"sid"` // しょぼかるIDの意味
	Title       string    `json:"title"`
	LastUpdated time.Time `json:"last_updated"`
	Episodes    []Episode `json:"episodes"`
	Info        `json:",inline"`
}

// Info ...
type Info struct {
	Songs []Song              `json:"songs"`
	Staff map[string][]string `json:"staff"`
	Cast  map[string][]string `json:"cast"`
}

// Episode ...
type Episode struct {
	Number string `json:"number"` // 1 | 1.5 | 1-3 | 総集編
	Title  string `json:"title"`
}

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
