package usecase

import (
	"math/rand"
	"time"
)

// Size of shorten URL
const size = 6

// Characters of shorten URL in alphanumeric
var runes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")


// Logic functiono to generate unique alias of shorten URL
func GenerateShortURL() string {

	// Seed for the randomizer to genereate random number everytime this function is called
	rand.Seed(time.Now().UnixNano())

	chars := make([]rune, size)

	for i := range chars {
		chars[i] = runes[rand.Intn(len(runes))]
	}

	return string(chars)
}