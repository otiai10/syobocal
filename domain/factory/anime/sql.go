package factory

import "animapi/domain/model/anime"
import "database/sql"

/**
 * ファクトリーは
 * モデルのコンストラクションの方法を隠蔽する
 * インフラから返ってきた値の形式や
 * 各モデルのコンストラクションに何が必要か
 * アプリケーションやレポジトリには見せない
 */
type AnimeFactory struct{}

func GetAnimeFactory() *AnimeFactory {
	return &AnimeFactory{}
}
func (f *AnimeFactory) FromRecord(record *sql.Row) *model.Anime {
	var id, title string
	_ = record.Scan(&id, &title)
	return &model.Anime{
		"0",
		title,
		"nil",
	}
}
