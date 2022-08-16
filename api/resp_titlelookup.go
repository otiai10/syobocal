package api

// TitleLookupResponse ...
// ex) http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20150315_000000-
type TitleLookupResponse struct {
	Result     Result `xml:"Result"`
	TitleItems struct {
		Items []TitleItem `xml:"TitleItem"`
	} `xml:"TitleItems"`
}

// TitleItem ...
type TitleItem struct {
	TID           int        `xml:"TID"`           // アニメタイトルID
	LastUpdate    SyoboiTime `xml:"LastUpdate"`    // 最近更新された時間
	Title         string     `xml:"Title"`         // アニメタイトル
	ShortTitle    string     `xml:"ShortTitle"`    // タイトルの省略形
	TitleYomi     string     `xml:"TitleYomi"`     // タイトルの日本語読み
	TitleEN       string     `xml:"TitleEN"`       // タイトルを英語にしたらどうなんの的なやつ（英語なければ""）
	TitleFlag     int        `xml:"TitleFlag"`     // たぶん再放送フラグ的なやつ
	Comment       string     `xml:"Comment"`       // コメント（？）なんかいろいろごちゃまぜ
	Category      Category   `xml:"Cat"`           // カテゴリ
	FirstYear     SyoboiInt  `xml:"FirstYear"`     // 最初に放送された年
	FirstMonth    SyoboiInt  `xml:"FirstMonth"`    // 最初に放送された月
	FirstEndYear  SyoboiInt  `xml:"FirstEndYear"`  // 最初の放送が終わった年（未終了なら0）
	FirstEndMonth SyoboiInt  `xml:"FirstEndMonth"` // 最初の放送が終わった月（未終了なら0）
	FirstChannel  string     `xml:"FirstCh"`       // 最初に放送された放送局
	Keywords      string     `xml:"Keywords"`      // 検索キーワード. カンマ区切り
	UserPoint     int        `xml:"UserPoint"`     // しょぼかるって投票できるらしいのでそのポイント
	UserPointRank int        `xml:"UserPointRank"` // しょぼかるって投票できるらしいのでその結果ランキング
	SubTitles     string     `xml:"SubTitles"`     // 各放映話のサブタイトル（改行区切り？）
}
