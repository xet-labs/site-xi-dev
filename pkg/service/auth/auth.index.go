package auth

import (
	"time"

	"xi/pkg/lib/cfg"

	"gorm.io/gorm"
)

type AuthService struct {
	JwtSecret    []byte // set from cfg.Api.JwtSecret
	AccessTTL    time.Duration
	RefreshTTL   time.Duration
	CookieDomain string
	CookieSecure bool

	Api *AuthApi
}

var Auth = &AuthService{
	JwtSecret:    []byte(cfg.Api.JwtSecret),
	AccessTTL:    15 * time.Minute,
	RefreshTTL:   7 * 24 * time.Hour,
	CookieDomain: cfg.Api.CookieDomain,  // set this in config; "" works too
	CookieSecure: cfg.Api.SecureCookies, // true in prod

	Api: Api,
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		JwtSecret:    []byte(cfg.Api.JwtSecret),
		AccessTTL:    15 * time.Minute,
		RefreshTTL:   7 * 24 * time.Hour,
		CookieDomain: cfg.Api.CookieDomain,  // set this in config; "" works too
		CookieSecure: cfg.Api.SecureCookies, // true in prod
	}
}
