package auth

import (
	"fmt"
	"net/http"
	"time"

	model_store "xi/internal/app/model/store"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a *AuthApi) Refresh(c *gin.Context) {
	cookieName := "refresh"
	raw, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	hashed := Auth.HashToken(raw)
	var rt model_store.RefreshToken
	if err := store.Db.Cli().Where("token_hash = ?", hashed).First(&rt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if rt.Revoked {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
		return
	}
	if time.Now().After(rt.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token expired"})
		return
	}

	// rotation: create new refresh token and mark old as replaced & revoked
	newRaw, newRec, err := Auth.GenRefreshTokenRecord(uint64(rt.UID), c.GetHeader("User-Agent"), c.ClientIP())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to rotate token"})
		return
	}

	// mark old token as revoked and set ReplacedById
	rt.Revoked = true
	rt.ReplacedById = &newRec.ID
	if err := store.Db.Cli().Save(&rt).Error; err != nil {
		// log but continue
	}

	// issue new access token for user
	uid := fmt.Sprintf("%d", rt.UID)
	access, err := Auth.GenAccessToken(uid, []string{"default"}, Auth.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	// set new cookie (raw token)
	maxAge := int(Auth.RefreshTTL.Seconds())
	c.SetCookie(cookieName, newRaw, maxAge, "/", Auth.CookieDomain, Auth.CookieSecure, true)

	// return access token and profile
	var user model_store.User
	if err := store.Db.Cli().First(&user, rt.UID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": access,
		"expires_in":   int(Auth.AccessTTL.Seconds()),
		"profile": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
			"avatar":   user.AvatarURL,
		},
	})
}
