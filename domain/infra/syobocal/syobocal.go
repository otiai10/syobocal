package infra

import "encoding/xml"
import "net/http"
import "io/ioutil"

//とりあえずここで参照するけど、循環したら考える
import "animapi/domain/model/syobocal"

// infra.SyobocalAPIはHTTPクライアントを所持してます
type SyobocalAPI struct {
	client ISyobocalHTTPClient
}
type ISyobocalHTTPClient interface {
	FindHoge() []byte
}
type SyobocalHTTPClient struct {
	baseURL string
}

func GetSyobocalAPI() SyobocalAPI {
	return SyobocalAPI{
		SyobocalHTTPClient{
			baseURL: "http://cal.syoboi.jp/db.php",
		},
	}
}

// func SyobocalApiOf(client SyobocalHTTPClient) SyobocalAPI {
func SyobocalApiOf(client ISyobocalHTTPClient) SyobocalAPI {
	return SyobocalAPI{
		client: client,
	}
}
func (api *SyobocalAPI) FindHoge() model.SyobocalResponse {
	body := api.client.FindHoge()
	return api.convert(body)
}

func (client SyobocalHTTPClient) FindHoge() []byte {

	url := client.baseURL + "?Command=TitleLookup&TID=*&LastUpdate=20140320_000000-"
	resp, e := http.Get(url)
	if e != nil {
		panic(e)
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}

	resp.Body.Close()
	return body
}
func (api *SyobocalAPI) convert(responseBody []byte) model.SyobocalResponse {

	responseRoot := model.SyobocalResponse{}

	e := xml.Unmarshal(responseBody, &responseRoot)
	if e != nil {
		panic(e)
	}

	return responseRoot
}
