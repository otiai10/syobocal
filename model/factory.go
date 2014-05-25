package model

import "github.com/otiai10/animapi/infrastructure"

func CreateProgramsFromSyobocalResponse(response infrastructure.SyobocalResponse) []Program {
	programs := []Program{}
	for _, item := range response.TitleItems.Items {
		if anime, e := CreateAnimeFromTitelItem(item); e == nil {
			programs = append(programs, CreateProgram(anime))
		}
	}
	return programs
}
func CreateProgram(anime Anime) Program {
	return Program{
		Anime: anime,
	}
}
func CreateAnimeFromTitelItem(item infrastructure.TitleItem) (anime Anime, e error) {
	anime = Anime{
		Title: item.Title,
	}
	return
}
