package auth

import "golang.org/x/crypto/bcrypt"

func (a *AuthPkg) HashPass(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

func (a *AuthPkg) CheckPass(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
