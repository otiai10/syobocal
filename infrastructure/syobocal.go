package infrastructure

import "fmt"
import "time"
import "net/http"
import "io/ioutil"

var (
	syobocal_base_url      = "http://cal.syoboi.jp/db.php"
	syobocal_request_query = "?Command=TitleLookup&TID=*&LastUpdate=%s-%s"
	syobocal_time_format   = "20060102_150405"
)

type syobocalRequestQuery struct {
	baseQuery string
}

func (q *syobocalRequestQuery) BuildBySince(since time.Duration) string {
	from := time.Now().Add(since).Format(syobocal_time_format)
	to := ""
	return q.baseQuery + fmt.Sprintf(syobocal_request_query, from, to)
}

type SyobocalHttpClient struct {
	baseURL string
	query   syobocalRequestQuery
}

func NewSyobocalClient() SyobocalHttpClient {
	return SyobocalHttpClient{
		baseURL: syobocal_base_url,
		query: syobocalRequestQuery{
			baseQuery: syobocal_request_query,
		},
	}
}
func (c *SyobocalHttpClient) buildURLBySince(since time.Duration) string {
	return c.baseURL + c.query.BuildBySince(since)
}
func (c *SyobocalHttpClient) TitleLookup(since time.Duration) (SyobocalResponse, error) {
	url := c.buildURLBySince(since)
	return c.requestSyobocal(url)
}
func (c *SyobocalHttpClient) requestSyobocal(url string) (sRes SyobocalResponse, e error) {
	resp, e := http.Get(url)
	if e != nil {
		return
	}
	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	return ConvertBytes2Response(bytes)
}
