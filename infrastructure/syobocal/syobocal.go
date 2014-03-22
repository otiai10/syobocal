package syobocal

import "encoding/xml"

// infra.SyobocalAPIはHTTPクライアントを所持してます
type SyobocalAPI struct {
	client ISyobocalHTTPClient
}

func SyobocalApiOf(client ISyobocalHTTPClient) SyobocalAPI {
	return SyobocalAPI{
		client: client,
	}
}
func GetAPI() SyobocalAPI {
	client := &SyobocalHTTPClient{
		baseURL: "http://cal.syoboi.jp/db.php",
	}
	return SyobocalAPI{client}
}
func (api *SyobocalAPI) convert(responseBody []byte) Response {
	responseRoot := Response{}
	e := xml.Unmarshal(responseBody, &responseRoot)
	if e != nil {
		panic(e)
	}
	return responseRoot
}

func (api *SyobocalAPI) RequestByRange(from, to string) Response {
	query := SyobocalQuery{
		Command: "TitleLookup",
		From:    from,
		To:      to,
	}
	body := api.client.ExecQuery(query)
	return api.convert(body)
}
