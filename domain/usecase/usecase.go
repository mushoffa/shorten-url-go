package usecase

import (
	"fmt"
	"errors"
	"regexp"
	"strings"
	"time"

	"shorten-url-go/domain/entity"
	"shorten-url-go/domain/repository"
)

var (
	ErrInvalidURL = errors.New("Invalid format, URL must begin with http or https")
)

type Usecase interface {
	EncodeURL(string) (string, error)
	DecodeURL(string) (string, error)
	GetURL(string) (entity.URL, error)
	GetAllURL() ([]entity.URL, error)
	ValidateURL(string) error
}

type usecase struct {
	repository domain.Repository
}

func NewUsecase(repository domain.Repository) Usecase {
	return &usecase{repository}
}

func (u *usecase) EncodeURL(originalURL string) (string, error) {
	if err := u.ValidateURL(originalURL); err != nil {
		return "", err
	}

	existingShortenURL, err := u.repository.FindByOriginalURL(originalURL)
	if err == nil {
		return existingShortenURL, nil
	}

	shortenURL := GenerateShortURL()

	now := time.Now().Local()
	timestamp := fmt.Sprintf("%d/%d/%d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())

	url := entity.URL {
		Created: timestamp,
		OriginalURL: originalURL,
		ShortenURL: shortenURL,
		RedirectCount: 0,
	}

	u.repository.AddURL(&url)

	return shortenURL, nil
}

func (u *usecase) DecodeURL(url string) (string, error) {
	data, err := u.repository.FindURL(url)
	if err != nil {
		return "", err
	}

	data.RedirectCount += 1

	return data.OriginalURL, nil
}

func (u *usecase) GetURL(url string) (entity.URL, error) {
	data, err := u.repository.FindURL(url)
	if err != nil {
		return entity.URL{}, err
	}

	return *data, nil
}

func (u *usecase) GetAllURL() ([]entity.URL, error) {
	return u.repository.FindAllURL()
}

func (u *usecase) ValidateURL(url string) error {
	regex, err := regexp.Compile("^(http|https)://")
	if err != nil {
		return err
	}

	url = strings.TrimSpace(url)
	if regex.MatchString(url) {
		return nil
	}

	return ErrInvalidURL
}