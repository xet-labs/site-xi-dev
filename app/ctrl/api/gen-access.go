package api

import (
    "time"
    
    "xi/app/lib/cfg"

    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID string   `json:"uid"`
    Scopes []string `json:"scopes"`
    jwt.RegisteredClaims
}

func GenAccessToken(userID string, scopes []string, ttl time.Duration) (string, error) {
    now := time.Now()
    claims := Claims{
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
    tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return tok.SignedString(cfg.Api.JwtSecret)
}
