package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string   `json:"uid"`
	Scopes []string `json:"scopes"`
	jwt.RegisteredClaims
}
