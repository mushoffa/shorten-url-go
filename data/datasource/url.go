package datasource

import (
	"shorten-url-go/domain/entity"
)

type URL interface {
	Set(string, *entity.URL) error
	Get(string) (*entity.URL, error)
	GetAll() ([]entity.URL, error)
	FindExistingURL(string) (string, error)
}