package controller_test

import (
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"shorten-url-go/data/controller"
	"shorten-url-go/data/model/controller"
	"shorten-url-go/domain/entity"
	"shorten-url-go/domain/usecase"
	"shorten-url-go/domain/repository"
	mocks "shorten-url-go/mocks/domain/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestEncode_Success(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	requestPayload := model.EncodeRequest{ URL: "www.test.com" }
	mockShortenURL := "4afX1F"
	mockResponsePayload := fmt.Sprintf("{\"shorten_url\":\"%v\"}", mockShortenURL)
	mockUsecase.On("EncodeURL").Return(mockShortenURL, nil)


	// When
	requestBody, _ := json.Marshal(requestPayload)
	request := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewReader(requestBody))
	request.Header.Add(`Content-Type`, `application/json`)
	response, _ := fiber.Test(request, -1)
	responseBody, _ := ioutil.ReadAll(response.Body)

	// Then
	utils.AssertEqual(t, 200, response.StatusCode)
	utils.AssertEqual(t, mockResponsePayload, string(responseBody))
}

func TestEncode_Payload_Error(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	// Empty payload

	// When
	request := httptest.NewRequest(http.MethodPost, "/encode", nil)
	request.Header.Add(`Content-Type`, `application/json`)
	response, _ := fiber.Test(request, -1)

	// Then
	utils.AssertEqual(t, 500, response.StatusCode)
}

func TestEncode_Usecase_Error(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	requestPayload := model.EncodeRequest{ URL: "www.test.com" }
	mockResponsePayload := fmt.Sprintf("{\"error\":\"%v\"}", usecase.ErrInvalidURL)
	mockUsecase.On("EncodeURL").Return("", usecase.ErrInvalidURL)


	// When
	requestBody, _ := json.Marshal(requestPayload)
	request := httptest.NewRequest(http.MethodPost, "/encode", bytes.NewReader(requestBody))
	request.Header.Add(`Content-Type`, `application/json`)
	response, _ := fiber.Test(request, -1)
	responseBody, _ := ioutil.ReadAll(response.Body)

	// Then
	utils.AssertEqual(t, 500, response.StatusCode)
	utils.AssertEqual(t, mockResponsePayload, string(responseBody))
}

func TestRedirect_Success(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	mockOriginalURL := "http://google.com"
	mockParam := "4afX1F"
	mockEndpoint := fmt.Sprintf("/r/%v",mockParam)
	mockUsecase.On("DecodeURL").Return(mockOriginalURL, nil)

	// When
	request := httptest.NewRequest(http.MethodGet, mockEndpoint, nil)
	response, _ := fiber.Test(request, -1)

	// Then
	utils.AssertEqual(t, 307, response.StatusCode)
	utils.AssertEqual(t, mockOriginalURL, response.Header["Location"][0])
}

func TestRedirect_Error(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	mockResponsePayload := fmt.Sprintf("{\"error\":\"%v\"}", domain.ErrURLNotFound)
	mockParam := "mock"
	mockEndpoint := fmt.Sprintf("/r/%v",mockParam)
	mockUsecase.On("DecodeURL").Return("", domain.ErrURLNotFound)


	// When
	request := httptest.NewRequest(http.MethodGet, mockEndpoint, nil)
	response, _ := fiber.Test(request, -1)
	responseBody, _ := ioutil.ReadAll(response.Body)

	// Then
	utils.AssertEqual(t, 500, response.StatusCode)
	utils.AssertEqual(t, mockResponsePayload, string(responseBody))
}

func TestFindByUrl_Success(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	mockData := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "wb2Hdx",
		RedirectCount: 0,
	}
	mockParam := "mock"
	mockEndpoint := fmt.Sprintf("/findByUrl/%v",mockParam)
	mockUsecase.On("GetURL").Return(mockData, nil)

	// When
	request := httptest.NewRequest(http.MethodGet, mockEndpoint, nil)
	response, _ := fiber.Test(request, -1)

	// Then
	utils.AssertEqual(t, 200, response.StatusCode)
}

func TestFindByUrl_Error(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	mockResponsePayload := fmt.Sprintf("{\"error\":\"%v\"}", domain.ErrURLNotFound)
	mockParam := "mock"
	mockEndpoint := fmt.Sprintf("/findByUrl/%v",mockParam)
	mockUsecase.On("GetURL").Return(entity.URL{}, domain.ErrURLNotFound)

	// When
	request := httptest.NewRequest(http.MethodGet, mockEndpoint, nil)
	response, _ := fiber.Test(request, -1)
	responseBody, _ := ioutil.ReadAll(response.Body)

	// Then
	utils.AssertEqual(t, 500, response.StatusCode)
	utils.AssertEqual(t, mockResponsePayload, string(responseBody))
}

func TestFindAllUrl_Success(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	mockData := entity.URL{
		Created: "14/11/2022 20:34:28",
		OriginalURL: "https://www.tokopedia.com/clavelluiofficial/clavellui-tas-selempang-dada-3-in-1-hp-dompet-card-organizer-unisex-excellent-navy?extParam=src%3Dmultiloc%26whid%3D11066452",
		ShortenURL: "wb2Hdx",
		RedirectCount: 0,
	}
	mockSlice := []entity.URL{mockData}
	mockUsecase.On("GetAllURL").Return(mockSlice, nil)

	// When
	request := httptest.NewRequest(http.MethodGet, "/findAll", nil)
	response, _ := fiber.Test(request, -1)

	// Then
	utils.AssertEqual(t, 200, response.StatusCode)
}

func TestFindAllUrl_Error(t *testing.T) {
	mockUsecase := new(mocks.MockUsecase)
	controller := controller.NewURLController(mockUsecase)

	fiber := fiber.New()
	controller.Router(fiber)

	// Given
	mockSlice := []entity.URL{}
	mockUsecase.On("GetAllURL").Return(mockSlice, domain.ErrURLNotFound)

	// When
	request := httptest.NewRequest(http.MethodGet, "/findAll", nil)
	response, _ := fiber.Test(request, -1)

	// Then
	utils.AssertEqual(t, 500, response.StatusCode)
}