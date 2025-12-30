package crypt

import (
	"crypto/rand"
)

// SecureRandom returns n bytes of cryptographically secure random.
func (c *CryptLib) Random(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}
