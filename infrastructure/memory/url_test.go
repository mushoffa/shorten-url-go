package memory_test

import (
	"testing"

	"shorten-url-go/domain/entity"
	"shorten-url-go/domain/repository"
	"shorten-url-go/infrastructure/memory"

	"github.com/stretchr/testify/assert"
)

func TestSet_Success(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// Given
	mockURL := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 0,
	}

	// When
	err := inmemory.Set(mockURL.ShortenURL, &mockURL)

	// Then
	assert.NoError(t, err)
}

func TestGet_Success(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// Given
	mockURL := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 0,
	}

	// When
	err := inmemory.Set(mockURL.ShortenURL, &mockURL)
	res, err := inmemory.Get(mockURL.ShortenURL)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, mockURL, *res)
}

func TestGet_Error(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// Given
	mockURL := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 0,
	}

	// When
	err := inmemory.Set(mockURL.ShortenURL, &mockURL)
	res, err := inmemory.Get("3afX1F")

	// Then
	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, domain.ErrURLNotFound, err)
}

func TestGetAll_Success(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// Given
	mockURL := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 1,
	}

	mockURL2 := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "wb2Hdx",
		RedirectCount: 3,
	}

	inmemory.Set(mockURL.ShortenURL, &mockURL)
	inmemory.Set(mockURL2.ShortenURL, &mockURL2)

	// When
	res, err := inmemory.GetAll()

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, mockURL, res[0])
	assert.Equal(t, mockURL2, res[1])
}

func TestGetAll_empty_Success(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// When
	res, err := inmemory.GetAll()

	// Then
	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestFindExistingURL_Success(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// Given
	mockURL := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 1,
	}
	inmemory.Set(mockURL.ShortenURL, &mockURL)

	// When
	res, err := inmemory.FindExistingURL(mockURL.OriginalURL)

	// Then
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestFindExistingURL_Error(t *testing.T) {
	inmemory := memory.NewURLMemory()

	// Given
	mockURL := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 1,
	}
	inmemory.Set(mockURL.ShortenURL, &mockURL)

	// When
	res, err := inmemory.FindExistingURL("http://www.otherurl.com")

	// Then
	assert.Error(t, err)	
	assert.Empty(t, res)
	assert.Equal(t, domain.ErrURLNotFound, err)
}