package util

import (
	"math/rand"
	"strings"
	"time"
)

// ==== CONSTANTS

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate n charactors long random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generate random owner name
func RandomOwner() string {
	return RandomString(8)
}

// Generate random money
func RandomMoney() int64 {
	return RandomInt(0, 1_00_000)
}

// Generate random currency
func RandomCurrency() string {
	currency := []string{"INR", "USD"}
	return currency[rand.Intn(len(currency))]
}
