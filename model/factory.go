package model

import "github.com/otiai10/animapi/infrastructure"

func CreateAnimeFromTitelItem(item infrastructure.TitleItem) (anime Anime, e error) {
	anisongs := CreateAnisongsFromCommentString(item.Comment, item.TID)
	anime = Anime{
		TID:      item.TID,
		Title:    item.Title,
		Anisongs: anisongs,
	}
	return
}
