package animapi

import "net/http"
import "time"

import "fmt"
import "io/ioutil"

type syobocal struct{}

var SYOBOCAL = syobocal{}

func (s syobocal) Greet() string {
	return "Hi, I'm Syobocal!"
}

// {{{
type Program struct {
	Title string
}

// }}}
func (s syobocal) FindPrograms(since time.Duration) /*[]Program*/ {
	url := "http://cal.syoboi.jp/db.php?Command=TitleLookup&TID=*&LastUpdate=20140525_000000-"
	resp, _ := http.Get(url)
	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("んごー %s\n", string(buf))
}
