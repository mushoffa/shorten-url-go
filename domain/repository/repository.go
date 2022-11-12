package domain

import (
	"shorten-url-go/domain/entity"
)

type Repository interface {
	AddURL(*entity.URL) error
	FindURL(string) (*entity.URL, error)
	FindAllURL() ([]entity.URL, error)
	FindByOriginalURL(string) (string, error)

	// GetShortenURL(string) (string, error)
	// GetOriginalURL(string) (string, error)
}