package model

import "strings"
import "errors"
import "regexp"

var (
	comment_section_separator = regexp.MustCompile("\n\n|\n\\*")
	comment_section_anisong   = regexp.MustCompile("(オープニングテーマ|エンディングテーマ|挿入歌)([0-9]*)「(.+)」$")
)

// その他のCommentSectionについても同様にしなきゃならん気がするんですがそれは
func CreateAnisongsFromCommentString(comment string, tid int) (anisongs []Anisong) {
	for _, sect := range comment_section_separator.Split(comment, -1) {
		if section, e := NewCommentSection(sect); e == nil {
			if is, label, index, title := section.isAnisongHeader(); is {
				anisongs = append(
					anisongs,
					NewAnisong(tid, title, label, index, section.Content),
				)
			}
		}
	}
	return
}

type CommentSection struct {
	raw     string
	Header  string
	Content string
}

func (s *CommentSection) RestoreHeaderAndContent() (e error) {
	// 先頭から改行以外の文字の連続がヘダー。それ以下はコンテンツ
	matches := strings.Split(s.raw, "\n")
	if len(matches) < 3 || matches[0] == "" {
		return errors.New("Invalid Comment Section")
	}
	s.Header = matches[0]
	s.Content = strings.Join(matches[1:], "\n")
	return
}
func NewCommentSection(sect string) (section CommentSection, e error) {
	section = CommentSection{raw: sect}
	if e = section.RestoreHeaderAndContent(); e != nil {
		return
	}
	return
}
func (s *CommentSection) isAnisongHeader() (is bool, label string, index string, title string) {
	matches := comment_section_anisong.FindStringSubmatch(s.Header)
	if len(matches) < 3 {
		return
	}
	if len(matches) == 4 {
		return true, matches[1], matches[2], matches[3]
	}
	return true, matches[1], "", matches[2]
}
