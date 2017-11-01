package anime

// Anime ...
type Anime struct {
	Title    string
	Episodes []Episode
	Songs    []Song
}

// Info ...
type Info struct {
	Songs []Song
	Staff map[string][]string
	Cast  map[string][]string
}

// Episode ...
type Episode struct {
	Number float32
	Title  string
}

// Song ...
type Song struct {
	Title    string
	Label    string   // オープニングとかエンディングとか挿入歌とか
	Words    []string // 作詞
	Music    []string // 作曲
	Composer []string // 編曲
	Singer   []string // 歌
}
