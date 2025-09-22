package auth

import (
	"net/http"
	"time"

	model_store "xi/internal/app/model/store"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
)

func (a *AuthApi) Logout(c *gin.Context) {
	cookieName := "refresh"
	raw, err := c.Cookie(cookieName)
	if err == nil {
		// revoke the token if present
		var rt model_store.RefreshToken
		db := store.Db.Cli()
		if err := db.Where("token_hash = ?", Auth.HashToken(raw)).First(&rt).Error; err == nil {
			rt.Revoked = true
			rt.UpdatedAt = time.Now()
			_ = db.Save(&rt)
		}
	}

	// clear cookie
	c.SetCookie(cookieName, "", -1, "/", Auth.CookieDomain, Auth.CookieSecure, true)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
