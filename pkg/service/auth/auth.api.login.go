package auth

import (
	"fmt"
	"net/http"
	"time"

	model_store "xi/internal/app/model/store"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginRequest shape
type LoginRequest struct {
	Email    string `json:"email" binding:"required,max=254"`
	Password string `json:"password" binding:"required,max=254"`
}

func (a *AuthApi) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request, " + err.Error()})
		return
	}

	var user model_store.User
	if err := store.Db.Cli().Where("username = ? OR email = ?", req.Email, req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email or username doesnt exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if user.Status != "active" {
		c.JSON(http.StatusForbidden, gin.H{"error": "account inactive"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credential"})
		return
	}

	// issue access token (string)
	access, err := Auth.GenAccessToken(fmt.Sprint(user.ID), []string{"default"}, Auth.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error, failed to create token"})
		return
	}

	// create refresh token record and set cookie
	ua := c.GetHeader("User-Agent")
	ip := c.ClientIP()
	rawRefresh, _, err := Auth.GenRefreshTokenRecord(user.ID, ua, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error, failed to create refresh token"})
		return
	}

	// set cookie. For security: HttpOnly, Secure, SameSite=Lax (or Strict if you prefer)
	maxAge := int(Auth.RefreshTTL.Seconds())
	cookieName := "refresh"
	c.SetCookie(cookieName, rawRefresh, maxAge, "/", Auth.CookieDomain, Auth.CookieSecure, true) // last true = HttpOnly
	// Note: Gin's SetCookie doesn't support setting SameSite directly; if you need SameSite, set header manually:
	// cookie := &http.Cookie{... SameSite: http.SameSiteLaxMode}
	// http.SetCookie(c.Writer, cookie)

	// update last_login time, optional
	now := time.Now()
	user.LastLogin = &now
	_ = store.Db.Cli().Save(&user)

	// Return access token and a small profile object (client stores access in memory)
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
