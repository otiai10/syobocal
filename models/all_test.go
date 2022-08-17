package models

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestAnime(t *testing.T) {
	Expect(t, Anime{}).TypeOf("models.Anime")
	Expect(t, Anime{}.Songs).TypeOf("[]models.Song")
}

func TestSong(t *testing.T) {
	Expect(t, Song{}).TypeOf("models.Song")
}
