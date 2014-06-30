package controllers

import "github.com/revel/revel"
import "github.com/otiai10/animapi"
import "github.com/otiai10/animapi/model"

type Anime struct {
	*revel.Controller
}

type Song struct {
	Anime model.Anime
	Song  model.Anisong
}

func (c Anime) Index() revel.Result {
	since, _ := animapi.Since("-12w")
	var songs []Song
	animes := animapi.DB("./my.conf", "test").FindAnimes(since)
  revel.INFO.Println("???", animes)
	for _, anime := range animes {
		for _, anisong := range animapi.DB("./my.conf", "test").FindAnisongsByTID(anime.TID) {
			song := Song{
				anime,
				anisong,
			}
			songs = append(songs, song)
		}
	}
	return c.Render(songs)
}
