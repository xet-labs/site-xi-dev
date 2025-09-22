package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"

	model_store "xi/internal/app/model/store"
	"xi/pkg/service/store"
)

// generateOpaqueToken generates a URL-safe random string
func GenOpaqueToken(nBytes int) (string, error) {
	b := make([]byte, nBytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// hashToken returns the SHA256 hex string of token
func (s *AuthService) HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

// Create & store refresh token record in DB, return raw token for client
func (s *AuthService) GenRefreshTokenRecord(uid uint64, ua, ip string) (rawToken string, record model_store.RefreshToken, err error) {
	raw, err := GenOpaqueToken(48)
	if err != nil {
		return "", record, err
	}
	rt := model_store.RefreshToken{
		UID:       uint(uid),
		Revoked:   false,
		TokenHash: s.HashToken(raw),
		UserAgent: &ua,
		IPAddress: &ip,
		ExpiresAt: time.Now().Add(s.RefreshTTL),
	}
	
	if err := store.Db.Cli().Create(&rt).Error; err != nil {
		return "", record, err
	}
	return raw, rt, nil
}
