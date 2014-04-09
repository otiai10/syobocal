package factory

import "animapi/domain/model/anime"
import "animapi/infrastructure/syobocal"

import "strconv"

/**
 * ファクトリーは
 * モデルのコンストラクションの方法を隠蔽する
 * インフラから返ってきた値の形式や
 * 各モデルのコンストラクションに何が必要か
 * アプリケーションやレポジトリには見せない
 */
func (f *AnimeFactory) FromSyobocalResponse(res syobocal.Response) []*model.Anime {
	var animes []*model.Anime
	for _, item := range res.TitleItems.TitleItem {
		animes = append(animes, f.item2anime(item))
	}
	return animes
}
func (f *AnimeFactory) item2anime(item syobocal.TitleItem) *model.Anime {
	fy, _ := strconv.Atoi(item.FirstYear)
	fm, _ := strconv.Atoi(item.FirstMonth)
	fey, _ := strconv.Atoi(item.FirstEndYear)
	fem, _ := strconv.Atoi(item.FirstEndMonth)
	return &model.Anime{
		TID:           item.TID,
		Title:         item.Title,
		Comment:       item.Comment,
		FirstYear:     fy,
		FirstMonth:    fm,
		FirstEndYear:  fey,
		FirstEndMonth: fem,
	}
}
