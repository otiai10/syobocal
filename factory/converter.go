package factory

import (
	"bytes"
	"regexp"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/otiai10/animapi/models"
	"github.com/otiai10/animapi/syobocal"
)

var (
	songHeader = regexp.MustCompile("^\\*(オープニングテーマ|エンディングテーマ|挿入歌)([0-9]*)「(.+)」")
	subtitle   = regexp.MustCompile("^\\*(.+)\\*(.+)")
)

// ConvertTitleLookupResponseToAnime syobocalパッケージのTitleLookupResponseをanimapi/modelsパッケージのAnimeに変換します.
func ConvertTitleLookupResponseToAnime(tlr syobocal.TitleLookupResponse) ([]*models.Anime, error) {
	animes := []*models.Anime{}
	// TODO: separate by items
	for _, item := range tlr.TitleItems.Items {
		anime := &models.Anime{
			ID:         item.TID,
			UpdatedAt:  item.LastUpdate.Value,
			Title:      item.Title,
			CommentRaw: item.Comment,
			Category:   models.Category(item.Category),
			Songs:      parseRawComment(item.TID, item.Comment),
			Programs:   parseRawSubTitles(item.TID, item.SubTitles),
			Keywords:   strings.Split(item.Keywords, ","),
			FirstBroadcast: time.Date(
				item.FirstYear.Value, time.Month(item.FirstMonth.Value), 1, 0, 0, 0, 0, time.UTC,
			),
			FirstEnded: mysql.NullTime{
				Time:  time.Date(item.FirstEndYear.Value, time.Month(item.FirstEndMonth.Value), 1, 0, 0, 0, 0, time.UTC),
				Valid: !(item.FirstEndYear.Value == 0 || item.FirstEndMonth.Value == 0),
			},
		}
		animes = append(animes, anime)
	}
	return animes, nil
}

// parseRawComment Comment
// フィールドを1行1行解釈して、songsを作っていく
// 将来的にはsongs以外も返す？
func parseRawComment(animeID int, raw string) []models.Song {
	rows := bytes.Split([]byte(raw), []byte("\n"))
	songs := []models.Song{}
	f := false
	for _, row := range rows {
		if len(row) == 0 {
			f = false
			continue
		}
		matches := songHeader.FindAllSubmatch(row, -1)
		if len(matches) < 1 {
			if f {
				keyValue := bytes.Split(bytes.Trim(row, ":"), []byte(":"))
				if len(keyValue) < 2 {
					continue
				}
				// 最後の要素のAttributesにぶちこむ
				songs[len(songs)-1].Attributes[string(keyValue[0])] = string(keyValue[1])
			}
			continue
		}
		f = true
		song := models.Song{
			AnimeID:    animeID,
			Type:       string(matches[0][1]),
			Number:     string(matches[0][2]),
			Title:      string(matches[0][3]),
			Attributes: map[string]string{},
		}
		songs = append(songs, song)
	}
	return songs
}

// parseRawSubTitles Subtitles
func parseRawSubTitles(animeID int, raw string) []models.Program {
	programs := []models.Program{}
	for _, row := range bytes.Split([]byte(raw), []byte("\n")) {
		matches := subtitle.FindAllSubmatch(row, -1)
		if len(matches) < 1 {
			continue
		}
		program := models.Program{
			AnimeID: animeID,
			Chapter: string(matches[0][1]),
			Title:   string(matches[0][2]),
		}
		programs = append(programs, program)
	}
	return programs
}
