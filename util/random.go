package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz" 
const capitalAlphabet = "ABCDEFGHIGKLMNOPQRSTUVWXYZ" 
var currencies = []string{"USD", "EUR", "CAD", "IR", "AUE", "PND"} 

func init()  {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates random nunber between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

//RandomString generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	// sb.WriteByte(capitalAlphabet[rand.Intn(k)])
	//remain name
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

//RandomCapitalString generates random string of capital characterd of length n
func RandomCapitalString(n int) string {
	var sb strings.Builder
	k := len(capitalAlphabet)

	// sb.WriteByte(capitalAlphabet[rand.Intn(k)])
	//remain name
	for i := 0; i < n; i++ {
		c := capitalAlphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomCapitalString(1) + RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1_000_000_000)
}

func RandomCurrency() string {
	l := len(currencies)

	return currencies[rand.Intn(l)]
}