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
	// if err := r.url.Set(data.ShortenURL, data); err != nil {
	// 	return err
	// }

	// if err := r.url.Set(data.OriginalURL, data); err != nil {
	// 	return err
	// }

	return r.url.Set(data.ShortenURL, data)
}

func (r *repository) FindURL(url string) (*entity.URL, error) {
	return r.url.Get(url)
}

func (r *repository) FindAllURL() ([]entity.URL, error) {
	return r.url.GetAll()
}

func (r *repository) FindByOriginalURL(url string) (string, error) {
	// data, err := r.url.Get(url)
	// if err != nil {
	// 	return "", err
	// }

	// return data.ShortenURL, nil
	return r.url.FindExistingURL(url)
}

// func (r *repository) GetShortenURL(url string) (string, error) {
// 	data, err := r.FindURL(url)
// 	if err != nil {
// 		return "", err
// 	}

// 	return data.ShortenURL, nil
// }

// func (r *repository) GetOriginalURL(url string) (string, error) {
// 	data, err := r.FindURL(url)
// 	if err != nil {
// 		return "", err
// 	}

// 	return data.OriginalURL, nil
// }