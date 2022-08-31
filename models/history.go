package models

import (
	"time"
)

type History struct {
	Context Context `json:"context"`
	Animes  []Anime `json:"animes"`
}

type Context struct {
	Time   time.Time `json:"time"`
	Query  string    `json:"query"`
	Output string    `json:"output"`
}
