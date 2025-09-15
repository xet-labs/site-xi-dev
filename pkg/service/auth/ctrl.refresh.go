package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model_db "xi/internal/app/model/db"
)

func (h *AuthCtrl) Refresh(c *gin.Context) {
	cookieName := "refresh"
	raw, err := c.Cookie(cookieName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	hashed := HashToken(raw) // helper we expose; or reimplement here
	var rec model_db.RefreshToken
	if err := h.S.DB.Where("refresh_token = ?", hashed).First(&rec).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if rec.Revoked {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token revoked"})
		return
	}
	if time.Now().After(rec.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token expired"})
		return
	}

	// rotation: create new refresh token and mark old as replaced & revoked
	ua := c.GetHeader("User-Agent")
	ip := c.ClientIP()
	newRaw, newRec, err := h.S.GenRefreshTokenRecord(uint64(rec.UID), ua, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to rotate token"})
		return
	}

	// mark old token as revoked and set ReplacedById
	rec.Revoked = true
	rec.ReplacedById = newRec.ID
	rec.UpdatedAt = time.Now()
	if err := h.S.DB.Save(&rec).Error; err != nil {
		// log but continue
	}

	// issue new access token for user
	uid := fmt.Sprintf("%d", rec.UID)
	access, err := h.S.GenAccessToken(uid, []string{"default"}, h.S.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	// set new cookie (raw token)
	maxAge := int(h.S.RefreshTTL.Seconds())
	c.SetCookie(cookieName, newRaw, maxAge, "/", h.S.CookieDomain, h.S.CookieSecure, true)

	// return access token and profile
	var user model_db.User
	if err := h.S.DB.First(&user, rec.UID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": access,
		"expires_in":   int(h.S.AccessTTL.Seconds()),
		"profile": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
			"avatar":   user.AvatarURL,
		},
	})
}
