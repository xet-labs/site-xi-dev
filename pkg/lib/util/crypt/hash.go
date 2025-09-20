package crypt

import "golang.org/x/crypto/bcrypt"

func (c *CryptLib) HashPass(pass string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return bytes, err
}

