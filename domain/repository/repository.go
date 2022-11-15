package domain

import (
	"errors"

	"shorten-url-go/domain/entity"
)

var (
	ErrURLNotFound = errors.New("URL not found on the system")
)

type Repository interface {
	AddURL(*entity.URL) error
	FindURL(string) (*entity.URL, error)
	FindAllURL() ([]entity.URL, error)
	FindByOriginalURL(string) (string, error)
}