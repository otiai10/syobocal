package model

type Anime struct {
	TID         int
	Title       string
	Anisongs    []Anisong
	LastUpdated int64
	Category    int
}
