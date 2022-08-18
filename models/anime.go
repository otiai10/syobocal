package models

import "time"

// Anime DBにEncode/Decodeしたいので、なるべくプリミティブがいい。
type Anime struct {
	TID         int       `json:"tid"` // しょぼかるにおけるTID
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
