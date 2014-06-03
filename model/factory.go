package model

import "github.com/otiai10/animapi/infrastructure"
import "time"

func CreateAnimeFromTitelItem(item infrastructure.TitleItem) (anime Anime, e error) {
	anisongs := CreateAnisongsFromCommentString(item.Comment, item.TID)
	ts, e := time.Parse(syobocal_time_format, item.LastUpdate)
	if e != nil {
		ts = time.Now()
	}
	anime = Anime{
		TID:         item.TID,
		Title:       item.Title,
		Anisongs:    anisongs,
		LastUpdated: ts.Unix(),
	}
	return
}
