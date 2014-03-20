package repo

// import "animapi/domain/model/anime"
import "animapi/domain/infra/syobocal"

// import "animapi/domain/factory/anime"
import "io/ioutil"
import "database/sql"
import "fmt"

// import "time"

// {{{ 旅行中はこっち使う
var fixturePath string = "../../../test/fixture/syobocal/syobocal.response.xml"

type DummyHTTPClient struct {
	*infra.SyobocalHTTPClient
	baseURL string
}

func (client DummyHTTPClient) FindHoge() []byte {
	var xml []byte
	xml, e := ioutil.ReadFile(fixturePath)
	if e != nil {
		panic(e)
	}
	return xml
}
func (client DummyHTTPClient) Query(query *SyobocalQuery) []byte {
	fmt.Println(
		"ほんとはこれを実行する",
		query.String(),
	)
	// {{{ とりあえず
	var xml []byte
	xml, e := ioutil.ReadFile(fixturePath)
	if e != nil {
		panic(e)
	}
	return xml
	// }}}
}
func (client DummyHTTPClient) FindOne(dsn, id string) (row *sql.Row) {
	return
}

// }}}
type SyobocalQuery struct {
	Command string
	/* うーん
	   From Time
	   To   Time
	*/
	From string
	To   string
}

func (q SyobocalQuery) String() string {
	return fmt.Sprintf(
		"?Command=%s&TID=*&LastUpdate=%s-%s",
		q.Command,
		q.From,
		q.To,
	)
}
