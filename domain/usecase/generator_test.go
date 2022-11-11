package usecase_test

import (
	"testing"

	"shorten-url-go/domain/usecase"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURL_Success(t *testing.T) {
	url := usecase.GenerateShortURL()
	assert.Equal(t, 6, len(url))
}

func TestGenerateRandomShortURL_Success(t *testing.T) {
	url1 := usecase.GenerateShortURL()
	url2 := usecase.GenerateShortURL()

	assert.True(t, (url1 != url2))
}