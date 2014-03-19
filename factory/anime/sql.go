package factory

import "animapi/model/anime"
import "database/sql"
/**
 * ファクトリーは
 * モデルのコンストラクションの方法を隠蔽する
 * インフラから返ってきた値の形式や
 * 各モデルのコンストラクションに何が必要か
 * アプリケーションやレポジトリには見せない
 */
func AnimeFromRecord(record *sql.Row) *model.Anime {
    var id, title string
    _ = record.Scan(&id, &title)
    return &model.Anime{
        title,
    }
}
