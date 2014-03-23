package factory

import "animapi/domain/model/anime"
import "animapi/infrastructure/syobocal"

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
	return &model.Anime{
		TID:     item.TID,
		Title:   item.Title,
		Comment: item.Comment,
	}
}
