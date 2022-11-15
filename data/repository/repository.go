package repository

import (
	"shorten-url-go/data/datasource"
	"shorten-url-go/domain/entity"
	"shorten-url-go/domain/repository"
)

type repository struct {
	url datasource.URL
}

func NewRepository(url datasource.URL) domain.Repository {
	return &repository{url}
}

func (r *repository) AddURL(data *entity.URL) error {
	return r.url.Set(data.ShortenURL, data)
}

func (r *repository) FindURL(url string) (*entity.URL, error) {
	return r.url.Get(url)
}

func (r *repository) FindAllURL() ([]entity.URL, error) {
	return r.url.GetAll()
}

func (r *repository) FindByOriginalURL(url string) (string, error) {
	return r.url.FindExistingURL(url)
}