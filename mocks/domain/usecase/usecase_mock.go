package mocks

import (
	"shorten-url-go/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) EncodeURL(url string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUsecase) DecodeURL(url string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUsecase) GetURL(url string) (entity.URL, error) {
	args := m.Called()
	return args.Get(0).(entity.URL), args.Error(1)
}

func (m *MockUsecase) GetAllURL() ([]entity.URL, error) {
	args := m.Called()
	return args.Get(0).([]entity.URL), args.Error(1)
}

func (m *MockUsecase) ValidateURL(url string) (error) {
	args := m.Called()
	return args.Error(1)
}