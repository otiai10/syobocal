package syobocal

import (
	"io/ioutil"
	"net/http"
)

type SyobocalHTTPClient struct {
	baseURL string
}

// コンストラクタ的なやつ
/* コンストラクタ的なやつ作りたいんだけど、なじぇ？
 * `not enough arguments in call to infra.SyobocalHTTPClient.New`
 * って叱られる
func (c SyobocalHTTPClient) New() *SyobocalHTTPClient {
	return &SyobocalHTTPClient{
		baseURL: "http://cal.syoboi.jp/db.php",
	}
}
*/
func (c *SyobocalHTTPClient) ExecQuery(query SyobocalQuery) (xml []byte) {

	url := c.baseURL + query.ToString()

	httpResponse, e := http.Get(url)
	if e != nil {
		panic(e)
	}

	xml, e = ioutil.ReadAll(httpResponse.Body)
	if e != nil {
		panic(e)
	}

	httpResponse.Body.Close()

	return xml
}
