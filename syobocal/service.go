package syobocal

import "net/http"

// Service なんか適当に
type Service interface {
	Get(string) (*http.Response, error)
}

// defaultService
type defaultService struct{}

func (s *defaultService) Get(uri string) (*http.Response, error) {
	return http.Get(uri)
}
