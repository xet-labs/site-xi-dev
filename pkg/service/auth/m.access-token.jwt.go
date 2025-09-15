package auth

import (
	"time"
	model_auth "xi/internal/app/model/auth"
	"xi/pkg/lib/cfg"

	"github.com/golang-jwt/jwt/v5"
)

func (s *AuthService) GenAccessToken(userID string, scopes []string, ttl time.Duration) (string, error) {
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
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.JwtSecret)
}

func (s *AuthService) ParseAccessToken(tokenStr string) (*model_auth.Claims, error) {
	tok, err := jwt.ParseWithClaims(tokenStr, &model_auth.Claims{}, func(t *jwt.Token) (any, error) {
		return s.JwtSecret, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	}
	if claims, ok := tok.Claims.(*model_auth.Claims); ok && tok.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
