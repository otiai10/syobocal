package anime

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// SectionHeaderExpression ...
	SectionHeaderExpression = regexp.MustCompile("^\\*[^\\*]+")

	// SongExpression ...
	SongExpression = regexp.MustCompile("^\\*(?P<label>(オープニング|エンディング|挿入歌))(テーマ)?[0-9]*「(?P<title>[^「^」]+)」")
	// SongMetaExpression ...
	SongMetaExpression = regexp.MustCompile("^:(?P<key>[^:]+):(?P<value>[^:]+)$")

	// StaffExpression ...
	StaffExpression = regexp.MustCompile("^\\*スタッフ")
	// StaffRoleExpression ...
	StaffRoleExpression = regexp.MustCompile("^:(?P<role>[^:]+):(?P<name>[^:]+)$")

	// CastExpression ...
	CastExpression = regexp.MustCompile("^\\*キャスト")
	// CastCharacterExpression ...
	CastCharacterExpression = regexp.MustCompile("^:(?P<chara>[^:]+):(?P<name>[^:]+)$")
)

// Token ...
type Token struct {
	Header string
	Lines  []string
}

// Type ...
func (t *Token) Type() string {
	fmt.Println(t.Header)
	if SongExpression.MatchString(t.Header) {
		return "song"
	}
	if StaffExpression.MatchString(t.Header) {
		return "staff"
	}
	if CastExpression.MatchString(t.Header) {
		return "cast"
	}
	return ""
}

// ParseComment ...
func ParseComment(raw string) (*Info, error) {
	info := new(Info)
	info.Staff = map[string][]string{}

	// Parse sections to Token
	tokens := []*Token{}
	var tmp *Token
	for _, line := range strings.Split(raw, "\n") {
		if SectionHeaderExpression.MatchString(line) {
			if tmp != nil {
				tokens = append(tokens, tmp)
			}
			tmp = new(Token)
			tmp.Header = line
			continue
		}
		if tmp != nil {
			tmp.Lines = append(tmp.Lines, line)
		}
	}
	if tmp != nil {
		tokens = append(tokens, tmp)
	}

	// Dispatch each Token
	for _, token := range tokens {
		switch token.Type() {
		case "song":
			info.Songs = append(info.Songs, CreateSong(token))
		case "staff":
			info.Staff = CreateStaff(token)
		case "cast":
			info.Cast = CreateCast(token)
		}
	}

	return info, nil
}

// CreateSong ...
func CreateSong(t *Token) Song {
	song := Song{}
	matched := SongExpression.FindAllStringSubmatch(t.Header, -1)
	for i, name := range SongExpression.SubexpNames() {
		switch name {
		case "label":
			song.Label = matched[0][i]
		case "title":
			song.Title = matched[0][i]
		}
	}
	for _, line := range t.Lines {
		match := SongMetaExpression.FindAllStringSubmatch(line, -1)
		if len(match) == 0 {
			continue
		}
		var key, val string
		for i, name := range SongMetaExpression.SubexpNames() {
			switch name {
			case "key":
				key = match[0][i]
			case "value":
				val = match[0][i]
			}
		}
		if val == "" {
			continue
		}
		switch key {
		case "作詞":
			song.Words = strings.Split(val, "、")
		case "作曲":
			song.Music = strings.Split(val, "、")
		case "編曲":
			song.Composer = strings.Split(val, "、")
		case "歌":
			song.Singer = strings.Split(val, "、")
		case "作詞・作曲・編曲":
			vals := strings.Split(val, "、")
			song.Words = vals
			song.Music = vals
			song.Composer = vals
		}
	}
	return song
}

// CreateStaff ...
func CreateStaff(t *Token) map[string][]string {
	staff := map[string][]string{}
	for _, line := range t.Lines {
		match := StaffRoleExpression.FindAllStringSubmatch(line, -1)
		if len(match) == 0 {
			continue
		}
		var role, name string
		for i, capture := range StaffRoleExpression.SubexpNames() {
			switch capture {
			case "role":
				role = match[0][i]
			case "name":
				name = match[0][i]
			}
		}
		staff[role] = append(staff[role], strings.Split(name, "、")...)
	}
	return staff
}

// CreateCast ...
func CreateCast(t *Token) map[string][]string {
	cast := map[string][]string{}
	for _, line := range t.Lines {
		match := CastCharacterExpression.FindAllStringSubmatch(line, -1)
		if len(match) == 0 {
			continue
		}
		var chara, name string
		for i, capture := range CastCharacterExpression.SubexpNames() {
			switch capture {
			case "chara":
				chara = match[0][i]
			case "name":
				name = match[0][i]
			}
		}
		cast[chara] = strings.Split(name, "、")
	}
	return cast
}
