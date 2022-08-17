package api

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Builder struct {
	command        Command
	lastUpdateFrom time.Time
	lastUpdateTo   time.Time
	fields         []Field
	tids           []int
}

type Command string

var (
	titleLookup Command = "TitleLookup"
	progLookup  Command = "ProgLookup"

	// TODO: 以下、まだ対応しなくていいや
	// chGroupLookup Command = "ChGroupLookup"
	// chLookup      Command = "ChLookup"
)

func TitleLookup() *Builder {
	return &Builder{
		command: titleLookup,
	}
}

func ProgLookup() *Builder {
	return &Builder{
		command: progLookup,
	}
}

// TODO: 以下、まだ対応しなくていいや
// func ChGroupLookup() *Builder {
// 	return &Builder{
// 		command: chGroupLookup,
// 	}
// }

// TODO: 以下、まだ対応しなくていいや
// func ChLookup() *Builder {
// 	return &Builder{
// 		command: chLookup,
// 	}
// }

func (b *Builder) LastUpdate(from, to time.Time) *Builder {
	b.lastUpdateFrom = from
	b.lastUpdateTo = to
	return b
}

func (b *Builder) Fields(fields ...Field) *Builder {
	b.fields = append(b.fields, fields...)
	return b
}

func (b *Builder) TID(tids ...int) *Builder {
	b.tids = append(b.tids, tids...)
	return b
}

func (b *Builder) Build() url.Values {

	query := url.Values{
		"Command": {string(b.command)},
	}

	// TIDの指定は、TitleLookupは指定無い場合 "*" が必須.
	// See https://docs.cal.syoboi.jp/spec/db.php/#%e3%82%bf%e3%82%a4%e3%83%88%e3%83%ab%e3%83%87%e3%83%bc%e3%82%bf%e3%81%ae%e5%8f%96%e5%be%97-titlelookup
	if len(b.tids) == 0 {
		if b.command == titleLookup {
			query["TID"] = []string{"*"}
		}
	} else {
		s := fmt.Sprintf("%d", b.tids[0])
		for _, tid := range b.tids[1:] {
			s += fmt.Sprintf(",%d", tid)
		}
	}

	// Fieldsは、指定が無ければqueryに含めない.
	// queryにFieldsが含まれない場合、AllFields扱いになる.
	// See https://docs.cal.syoboi.jp/spec/db.php/#%e3%82%bf%e3%82%a4%e3%83%88%e3%83%ab%e3%83%87%e3%83%bc%e3%82%bf%e3%81%ae%e5%8f%96%e5%be%97-titlelookup
	if len(b.fields) != 0 {
		query["Fields"] = []string{
			strings.Join(fieldsToStringSlice(b.fields), ","),
		}
	}

	// LastUpdateの始点・終点の指定が特に無い場合は、0時間を入れる.
	// See https://docs.cal.syoboi.jp/spec/db.php/#%e3%82%bf%e3%82%a4%e3%83%88%e3%83%ab%e3%83%87%e3%83%bc%e3%82%bf%e3%81%ae%e5%8f%96%e5%be%97-titlelookup
	var lastupdate string
	if !b.lastUpdateFrom.IsZero() {
		lastupdate += b.lastUpdateFrom.Format(queryTimeFormat)
	}
	lastupdate += "-"
	if !b.lastUpdateTo.IsZero() {
		lastupdate += b.lastUpdateTo.Format(queryTimeFormat)
	}
	if lastupdate != "-" {
		query["LastUpdate"] = []string{lastupdate}
	}

	return query
}
