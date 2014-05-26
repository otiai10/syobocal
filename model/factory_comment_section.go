package model

import "fmt"

import "strings"
import "errors"
import "regexp"

var (
	comment_section_separator = regexp.MustCompile("\n\n")
)

func CreateAnisongsFromCommentString(comment string) /* []Anisong */ {
	for _, sect := range comment_section_separator.Split(comment, -1) {
		_, _ = NewCommentSection(sect)
	}
}

type CommentSection struct {
	raw     string
	Header  string
	Content string
}

func (s *CommentSection) RestoreHeaderAndContent() (e error) {
	// 先頭から改行以外の文字の連続がヘダー。それ以下はコンテンツ
	matches := strings.Split(s.raw, "\n")
	if len(matches) < 3 {
		return errors.New("Invalid Comment Section")
	}
	s.Header = matches[1]
	s.Content = strings.Join(matches[2:], "\n")
	fmt.Printf("コンテンツ=======\n%+v\n\n", s.Content)
	return
}
func NewCommentSection(sect string) (section CommentSection, e error) {
	section = CommentSection{raw: sect}
	if e = section.RestoreHeaderAndContent(); e != nil {
		return
	}
	return
}
