package usecase

import (
	"math/rand"
	"time"
)

const size = 6

var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	chars := make([]rune, size)

	for i := range chars {
		chars[i] = runes[rand.Intn(len(runes))]
	}

	return string(chars)
}