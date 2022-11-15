package memory

import (
	"shorten-url-go/data/datasource"
	"shorten-url-go/domain/entity"
	"shorten-url-go/domain/repository"
)

type url struct {
	data map[string] *entity.URL
}

func NewURLMemory() datasource.URL {
	return &url{data: make(map[string]*entity.URL)}
}

func (m *url) Set(id string, url *entity.URL) error {
	m.data[id] = url

	return nil
}

func (m *url) Get(id string) (*entity.URL, error) {
	if url, exist := m.data[id]; exist {
		return url, nil
	}

	return nil, domain.ErrURLNotFound
}

func (m *url) GetAll() ([]entity.URL, error) {
	var urls []entity.URL

	for _, url := range m.data {
		urls = append(urls, *url)
	}

	return urls, nil
}

func (m *url) FindExistingURL(url string) (string, error) {

	for _, data := range m.data {
		if url == data.OriginalURL {
			return data.ShortenURL, nil
		}
	}

	return "", domain.ErrURLNotFound
}