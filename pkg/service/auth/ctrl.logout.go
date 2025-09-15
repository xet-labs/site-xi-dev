package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	model_db "xi/internal/app/model/db"
)

func (h *AuthCtrl) Logout(c *gin.Context) {
	cookieName := "refresh"
	raw, err := c.Cookie(cookieName)
	if err == nil {
		hashed := HashToken(raw)
		// revoke the token if present
		var rec model_db.RefreshToken
		if err := h.S.DB.Where("refresh_token = ?", hashed).First(&rec).Error; err == nil {
			rec.Revoked = true
			rec.UpdatedAt = time.Now()
			_ = h.S.DB.Save(&rec)
		}
	}

	// clear cookie
	c.SetCookie(cookieName, "", -1, "/", h.S.CookieDomain, h.S.CookieSecure, true)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
