package api

type Field string

var (
	TID           Field = "TID"
	LastUpdate    Field = "LastUpdate"
	Title         Field = "Title"
	ShortTitle    Field = "ShortTitle"
	TitleYomi     Field = "TitleYomi"
	TitleEN       Field = "TitleEN"
	TitleFlag     Field = "TitleFlag"
	Comment       Field = "Comment"
	Cat           Field = "Cat"
	FirstYear     Field = "FirstYear"
	FirstMonth    Field = "FirstMonth"
	FirstEndYear  Field = "FirstEndYear"
	FirstEndMonth Field = "FirstEndMonth"
	FirstChannel  Field = "FirstChannel"
	Keywords      Field = "Keywords"
	UserPoint     Field = "UserPoint"
	UserPointRank Field = "UserPointRank"
	SubTitles     Field = "SubTitles"
)

var (
	AllFields = []Field{
		TID,
		LastUpdate,
		Title,
		ShortTitle,
		TitleYomi,
		TitleEN,
		TitleFlag,
		Comment,
		Cat,
		FirstYear,
		FirstMonth,
		FirstEndYear,
		FirstEndMonth,
		FirstChannel,
		Keywords,
		UserPoint,
		UserPointRank,
		SubTitles,
	}
)

func fieldsToStringSlice(fields []Field) (s []string) {
	for _, f := range fields {
		s = append(s, string(f))
	}
	return s
}
