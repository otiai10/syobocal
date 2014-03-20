package infra

import (
	"fmt"
)

type SyobocalQuery struct {
	Command string
	From    string
	To      string //Timeの方がいいかな？
}

func (q SyobocalQuery) ToString() string {
	return fmt.Sprintf(
		"?Command=%s&TID=*&LastUpdate=%s-%s",
		q.Command,
		q.From,
		q.To,
	)
}
