package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	model_db "xi/internal/app/model/db"
)

func (a *AuthApi) Logout(c *gin.Context) {
	cookieName := "refresh"
	raw, err := c.Cookie(cookieName)
	if err == nil {
		hashed := HashToken(raw)
		// revoke the token if present
		var rec model_db.RefreshToken
		if err := Auth.DB.Where("refresh_token = ?", hashed).First(&rec).Error; err == nil {
			rec.Revoked = true
			rec.UpdatedAt = time.Now()
			_ = Auth.DB.Save(&rec)
		}
	}

	// clear cookie
	c.SetCookie(cookieName, "", -1, "/", Auth.CookieDomain, Auth.CookieSecure, true)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
