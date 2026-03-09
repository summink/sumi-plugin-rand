package random

import (
	"math/rand/v2"
	"strings"
)

const defaultCharsets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomChar(charsets string) string {
	return string(charsets[rand.IntN(len(charsets))])
}

func RandomChar() string {
	return randomChar(defaultCharsets)
}

func RandomString(minLen int, maxLen int) string {
	return RandomStringn(minLen, maxLen, defaultCharsets)
}

func RandomStringn(minLen int, maxLen int, charsets string) string {
	var value strings.Builder

	for range RandomIntn(minLen, maxLen) {
		value.WriteString(randomChar(charsets))
	}

	return value.String()
}
