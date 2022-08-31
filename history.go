package main

import (
	"time"

	"github.com/otiai10/syobocal/models"
)

type History struct {
	Context Context        `json:"context"`
	Animes  []models.Anime `json:"animes"`
}

type Context struct {
	Time   time.Time `json:"time"`
	Query  string    `json:"query"`
	Output string    `json:"output"`
}
