package auth

import (
	"errors"
	"time"

	"gorm.io/gorm"

	// "xi/pkg/lib"
	"xi/pkg/lib/cfg"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenNotFound      = errors.New("refresh token not found")
	ErrTokenRevoked       = errors.New("refresh token revoked")
	ErrTokenExpired       = errors.New("refresh token expired")
)

type AuthService struct {
	DB           *gorm.DB
	JwtSecret    []byte // set from cfg.Api.JwtSecret
	AccessTTL    time.Duration
	RefreshTTL   time.Duration
	CookieDomain string
	CookieSecure bool

	Api *AuthApi
}

var Auth = &AuthService{
	// DB:           lib.Db.Cli,
	JwtSecret:    []byte(cfg.Api.JwtSecret),
	AccessTTL:    15 * time.Minute,
	RefreshTTL:   7 * 24 * time.Hour,
	CookieDomain: cfg.Api.CookieDomain,  // set this in config; "" works too
	CookieSecure: cfg.Api.SecureCookies, // true in prod

	Api: Api, 
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB:           db,
		JwtSecret:    []byte(cfg.Api.JwtSecret),
		AccessTTL:    15 * time.Minute,
		RefreshTTL:   7 * 24 * time.Hour,
		CookieDomain: cfg.Api.CookieDomain,  // set this in config; "" works too
		CookieSecure: cfg.Api.SecureCookies, // true in prod
	}
}
