package repository_test

import (
	"testing"

	"shorten-url-go/data/repository"
	"shorten-url-go/domain/entity"

	mocks "shorten-url-go/mocks/data/datasource"

	"github.com/stretchr/testify/assert"
)

func TestAddURL_Success(t *testing.T) {
	mockDatasource := new(mocks.MockURLDatasource)
	repository := repository.NewRepository(mockDatasource)

	// Given
	mockData := entity.URL{}
	mockDatasource.On("Set").Return(nil)

	// When
	err := repository.AddURL(&mockData)

	// Then
	assert.NoError(t, err)
}

func TestFindURL_Success(t *testing.T) {
	mockDatasource := new(mocks.MockURLDatasource)
	repository := repository.NewRepository(mockDatasource)

	// Given
	mockData := &entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 3,
	}
	mockDatasource.On("Get").Return(mockData, nil)

	// When
	res, err := repository.FindURL("mockURL")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, mockData, res)
}

func TestFindAllURL_empty_Success(t *testing.T) {
	mockDatasource := new(mocks.MockURLDatasource)
	repository := repository.NewRepository(mockDatasource)

	// Given
	mockURLSlice := []entity.URL{}
	mockDatasource.On("GetAll").Return(mockURLSlice, nil)

	// When
	res, err := repository.FindAllURL()

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
	assert.Empty(t, res)
}

func TestFindByOriginalURL_Success(t *testing.T) {
	mockDatasource := new(mocks.MockURLDatasource)
	repository := repository.NewRepository(mockDatasource)

	// Given
	mockShortenURL := "4afX1F"
	mockDatasource.On("FindExistingURL").Return(mockShortenURL, nil)

	// When
	res, err := repository.FindByOriginalURL("mockURL")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, mockShortenURL, res)
}