package model

import "strconv"

type Anisong struct {
	TID    int
	Title  string
	Label  string
	Index  int
	Detail string
}

func NewAnisong(tid int, title, label, index, detail string) Anisong {
	i, _ := strconv.Atoi(index)
	return Anisong{
		TID:    tid,
		Title:  title,
		Label:  label,
		Index:  i,
		Detail: detail,
	}
}
