package model

// import "fmt"
// import "strings"
import "errors"
import "regexp"

var (
	comment_section_separator = regexp.MustCompile("\n\n")
	// 先頭から改行以外の文字の連続がヘダー。それ以下はコンテンツ
	comment_section_expresson = regexp.MustCompile("^([^\n]+)\n(.*)")
)

func CreateAnisongsFromCommentString(comment string) /* []Anisong */ {
	for _, sect := range comment_section_separator.Split(comment, -1) {
		// section := NewCommentSection(sect)
		_, _ = NewCommentSection(sect)
	}
}

type CommentSection struct {
	raw     string
	Header  string
	Content string
}

func (s *CommentSection) RestoreHeaderAndContent() (e error) {
	matches := comment_section_expresson.FindStringSubmatch(s.raw)
	if len(matches) < 2 {
		return errors.New("セクションとして妥当じゃない")
	}
	return
}
func NewCommentSection(sect string) (section CommentSection, e error) {
	section = CommentSection{raw: sect}
	e = section.RestoreHeaderAndContent()
	return
}

/*
func CreateAnisongFromBlockString(block string) Anisong  {
    header, value := c.extractHeaderAndValue(block)
    anisong
}
func (c *Comment) ToAnisongs() {
    c.splitBlocks()
    for _, bl := range c.blocks {
        header, value := c.extractHeaderAndValue(block)
        if ! c.isSongHeader(header) {
            continue
        }
        c.anisongs = append(c.anisongs, CreateAnisongFromBlockString(bl))
    }
}
func (c *Comment) splitBlocks() {
	re := regexp.MustCompile("[\n]{2}")
	c.blocks = re.Split(c.raw, -1)
}
func (c *Comment) extractAnisongs(block string) {
    header, value := c.extractHeaderAndValue(block)
    CreateAnisongFromStrings(header, value)
    c.anisongs = append(c.anisongs, Anisong{c.getSongTitle(header)})
}
func (c *Comment) extractHeaderAndValue(block string) (header, value string) {
	_r := regexp.MustCompile("^([^\n]+)\n(.*)")
	submatch := _r.FindSubmatch([]byte(block))
    if len(submatch) < 2 {
        return
    }
	header = string(submatch[1])
	rplcr := strings.NewReplacer(header+"\n", "")
	value = rplcr.Replace(block)
	return header, value
}
func (c *Comment) isSongHeader(header string) bool {
	re := regexp.MustCompile(`オープニングテーマ`)
	if re.Match([]byte(header)) {
		return true
	}
	re = regexp.MustCompile(`エンディングテーマ`)
	if re.Match([]byte(header)) {
		return true
	}
	re = regexp.MustCompile(`挿入歌`)
	if re.Match([]byte(header)) {
		return true
	}
	return false
}
func (c *Comment) getSongTitle(header string) (title string) {
	re := regexp.MustCompile("「(.*)」$")
	matche := re.FindSubmatch([]byte(header))
	return string(matche[1])
}
*/
