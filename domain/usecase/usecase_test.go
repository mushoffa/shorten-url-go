package usecase_test

import (
	"fmt"
	"errors"
	"testing"

	"shorten-url-go/domain/entity"
	"shorten-url-go/domain/repository"
	"shorten-url-go/domain/usecase"

	mocks "shorten-url-go/mocks/domain/repository"

	"github.com/stretchr/testify/assert"
)

func TestEncodeURL_Success(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockURL := "http://www.hypefast.id"
	mockRepository.On("FindByOriginalURL").Return("", errors.New("Mock error"))
	mockRepository.On("AddURL").Return(nil)

	// When
	res, err := useCase.EncodeURL(mockURL)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 6, len(res))
	assert.NotEqual(t, mockURL, res)
}

func TestEncodeURL_existing_Success(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockURL := "http://www.hypefast.id"
	mockExistingURL := "Pe3LV1"
	mockRepository.On("FindByOriginalURL").Return(mockExistingURL, nil)

	// When
	res, err := useCase.EncodeURL(mockURL)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 6, len(res))
	assert.Equal(t, mockExistingURL, res)
}

func TestEncodeURL_Error(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockURL := "www.google.com"

	// When
	res, err := useCase.EncodeURL(mockURL)

	// Then
	assert.Error(t, err)
	assert.Equal(t, "", res)
}

func TestDecodeURL_Success(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockData := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "4afX1F",
		RedirectCount: 0,
	}
	mockRepository.On("FindURL").Return(&mockData, nil)

	// When
	res, err := useCase.DecodeURL("4afX1F")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, mockData.OriginalURL, res)
	assert.Equal(t, uint32(1), mockData.RedirectCount)
}

func TestDecodeURL_Error(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockRepository.On("FindURL").Return(&entity.URL{}, domain.ErrURLNotFound)

	// When
	res, err := useCase.DecodeURL("4afX1F")

	// Then
	assert.Error(t, err)
	assert.Equal(t, "", res)
	assert.Equal(t, domain.ErrURLNotFound, err)
}

func TestGetURL_Success(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockData := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "wb2Hdx",
		RedirectCount: 0,
	}
	mockRepository.On("FindURL").Return(&mockData, nil)

	// When
	res, err := useCase.GetURL("wb2Hdx")

	// Then
	assert.NoError(t, err)
	assert.Equal(t, mockData, res)
}

func TestGetURL_Error(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockRepository.On("FindURL").Return(&entity.URL{}, domain.ErrURLNotFound)

	// When
	_, err := useCase.GetURL("wb2Hdx")

	// Then
	assert.Error(t, err)
	assert.Equal(t, domain.ErrURLNotFound, err)
}

func TestGetAllURL_empty_Success(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	mockURLSlice := []entity.URL{}
	mockRepository.On("FindAllURL").Return(mockURLSlice, nil)

	// When
	res, err := useCase.GetAllURL()

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 0, len(res))
}

func TestValidateURL_Success(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	testsValidURL := []struct {
		in string
		out error
	}{
		{
			in: "http://hypefast.id",
			out: nil,
		},
		{
			in: "https://hypefast.id",
			out: nil,
		},
		{
			in: "http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
			out: nil,
		},
		{
			in: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
			out: nil,
		},
	}

	for i, _ := range testsValidURL {
		tc := testsValidURL[i]
		t.Run(fmt.Sprintf("Test: %v", tc.in), func(t *testing.T) {
			t.Parallel()

			// When
			err := useCase.ValidateURL(tc.in)

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}

func TestValidateURL_Error(t *testing.T) {
	mockRepository := new(mocks.MockRepository)
	useCase := usecase.NewUsecase(mockRepository)

	// Given
	testsInvalidURL := []struct {
		in string
		out error
	}{
		{
			in: "www.hypefast.id",
			out: usecase.ErrInvalidURL,
		},
		{
			in: "www.hypefast.id",
			out: usecase.ErrInvalidURL,
		},
		{
			in: "www.chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity|Chart",
			out: usecase.ErrInvalidURL,
		},
		{
			in: "www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
			out: usecase.ErrInvalidURL,
		},
	}

	for i, _ := range testsInvalidURL {
		tc := testsInvalidURL[i]
		t.Run(fmt.Sprintf("Test: %v", tc.in), func(t *testing.T) {
			t.Parallel()

			// When
			err := useCase.ValidateURL(tc.in)

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}