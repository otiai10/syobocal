package infra

import "encoding/xml"
import "net/http"
import "io/ioutil"
import "animapi/model/syobocal"//とりあえずここで参照するけど、循環したら考える

import "fmt"

type SyobocalClient struct {
    baseURL string
}

func GetSyobocalAPI() SyobocalClient {
    return SyobocalClient{
        baseURL: "http://cal.syoboi.jp/db.php",
    }
}
func (c *SyobocalClient) Hoge() {
    // TODO: エラーハンドリング雑だなー
    url := c.baseURL + "?Command=TitleLookup&TID=*&LastUpdate=20140320_000000-"
    resp, _ := http.Get(url)
    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Println(
        c.convert(body),
    )
}

func (c *SyobocalClient)convert(responseBody []byte) model.SyobocalResponse {
    responseRoot := model.SyobocalResponse{}
    e := xml.Unmarshal(responseBody, &responseRoot)
    if e != nil {
         panic(e)
    }
    return responseRoot
}
