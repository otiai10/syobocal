package syobocal

type ISyobocalHTTPClient interface {
	ExecQuery(query SyobocalQuery) []byte
}
