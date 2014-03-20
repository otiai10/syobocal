package infra

import "encoding/xml"

//とりあえずここで参照するけど、循環したら考える
import "animapi/domain/model/syobocal"

// infra.SyobocalAPIはHTTPクライアントを所持してます
type SyobocalAPI struct {
	client ISyobocalHTTPClient
}

func SyobocalApiOf(client ISyobocalHTTPClient) SyobocalAPI {
	return SyobocalAPI{
		client: client,
	}
}
func (api *SyobocalAPI) convert(responseBody []byte) model.SyobocalResponse {
	responseRoot := model.SyobocalResponse{}
	e := xml.Unmarshal(responseBody, &responseRoot)
	if e != nil {
		panic(e)
	}
	return responseRoot
}

func (api *SyobocalAPI) RequestQuery(from, to string) model.SyobocalResponse {
	query := SyobocalQuery{
		Command: "TitleLookup",
		From:    from,
		To:      to,
	}
	body := api.client.ExecQuery(query)
	return api.convert(body)
}
