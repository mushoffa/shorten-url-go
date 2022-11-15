package mocks

import (
	"shorten-url-go/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockURLDatasource struct  {
	mock.Mock
}

func (m *MockURLDatasource) Set(id string, data *entity.URL) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockURLDatasource) Get(id string) (*entity.URL, error) {
	args := m.Called()
	return args.Get(0).(*entity.URL), args.Error(1)
}

func (m *MockURLDatasource) GetAll() ([]entity.URL, error) {
	args := m.Called()
	return args.Get(0).([]entity.URL), args.Error(1)
}

func (m *MockURLDatasource) FindExistingURL(url string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}