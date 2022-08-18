package factory

import (
	"time"

	"github.com/otiai10/syobocal/api"
	"github.com/otiai10/syobocal/models"
)

func ToAnimeListFromTitleLookup(res api.TitleLookupResponse) ([]models.Anime, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	animes := []models.Anime{}
	for _, item := range res.TitleItems.Items {
		anime := models.Anime{
			TID:   item.TID,
			Title: item.Title,
		}
		info, err := ParseComment(item.Comment, anime)
		if err != nil {
			return nil, err
		}
		anime.Songs = info.Songs
		anime.Cast = info.Cast
		anime.Staff = info.Staff
		anime.LastUpdated = time.Time(item.LastUpdate).In(loc)
		animes = append(animes, anime)
	}
	return animes, nil
}
