package begenerate

import (
	"crypto/rand"
	"math/big"
)

const defaultRandomAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// PseudoIdString generates a cryptographically random string with the specified length.
//
// The generated string matches [A-Za-z0-9]+ and it's transparent to URL-encoding.
func PseudoIdString(length int) string {
	return PseudoIdStringWithAlphabet(length, defaultRandomAlphabet)
}

// PseudoIdStringWithAlphabet generates a cryptographically random string
// with the specified length and characters set.
func PseudoIdStringWithAlphabet(length int, alphabet string) string {
	b := make([]byte, length)
	max := big.NewInt(int64(len(alphabet)))

	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		b[i] = alphabet[n.Int64()]
	}

	return string(b)
}

// PseudoIdRandomString generates a cryptographically random string with the specified length.
//
// This function uses crypto/rand for secure randomness, similar to RandomString.
// It's recommended to use RandomString instead of PseudoIdRandomString.
func PseudoIdRandomString() string {
	return PseudoIdString(16)
}
