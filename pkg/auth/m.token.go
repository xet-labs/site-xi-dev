package auth

import (
	"time"

	model_auth "xi/internal/app/model/auth"
	// model_db "xi/internal/app/model/db"
	"xi/pkg/lib/cfg"

	"github.com/golang-jwt/jwt/v5"
)

func (a *AuthPkg) GenAccessToken(userID string, scopes []string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := model_auth.Claims{
		UserID: userID,
		Scopes: scopes,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    cfg.Org.Name,
			Subject:   userID,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(cfg.Api.JwtSecret)
}

func (a *AuthPkg) GenRefreshToken(userID string, ttl time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   userID,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(cfg.Api.JwtSecret)
}

func (a *AuthPkg) ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return cfg.Api.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}