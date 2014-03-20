package infra

type ISyobocalHTTPClient interface {
	ExecQuery(query SyobocalQuery) []byte
}
