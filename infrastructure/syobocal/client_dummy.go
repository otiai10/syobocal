package infra

import (
	"fmt"
	"io/ioutil"
)

type SyobocalDummyHTTPClient struct {
	baseURL string
}

func (c *SyobocalDummyHTTPClient) ExecQuery(query SyobocalQuery) (xml []byte) {

	url := c.baseURL + query.ToString()
	fmt.Println("このリクエストをスタブする", url)
	// httpResponse, e := http.Get(url)
	// if e != nil { panic(e) }
	// xml, e = ioutil.ReadAll(httpResponse.Body)
	// if e != nil { panic(e) }
	// httpResponse.Body.Close()
	fixturePath := "../../../test/fixture/syobocal/syobocal.response.xml"
	xml, e := ioutil.ReadFile(fixturePath)
	if e != nil {
		panic(e)
	}

	return xml
}
