package str

import (
	"math/rand"
	"time"
)

// RandAlphanumericString ...
func RandAlphanumericString(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return StringWithCharset(length, charset)
}

// RandLowerAlphanumericString ...
func RandLowerAlphanumericString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	return StringWithCharset(length, charset)
}

// StringWithCharset ...
func StringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomNumericString ...
func RandomNumericString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(48, 57))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
