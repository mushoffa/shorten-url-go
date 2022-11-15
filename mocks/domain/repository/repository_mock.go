package mocks

import (
	"shorten-url-go/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) AddURL(data *entity.URL) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) FindURL(url string) (*entity.URL, error) {
	args := m.Called()
	return args.Get(0).(*entity.URL), args.Error(1)
}

func (m *MockRepository) FindAllURL() ([]entity.URL, error) {
	args := m.Called()
	return args.Get(0).([]entity.URL), args.Error(1)
}

func (m *MockRepository) FindByOriginalURL(url string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}