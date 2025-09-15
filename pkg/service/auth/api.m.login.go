package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	model_db "xi/internal/app/model/db"
)

// LoginRequest shape
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *AuthApi) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user model_db.User
	if err := Auth.DB.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if user.Status != "active" {
		c.JSON(http.StatusForbidden, gin.H{"error": "account not active"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		// optionally increment failed login counter here for lockout
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// issue access token (string)
	access, err := Auth.GenAccessToken(fmt.Sprint(user.ID), []string{"default"}, Auth.AccessTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	// create refresh token record and set cookie
	ua := c.GetHeader("User-Agent")
	ip := c.ClientIP()
	rawRefresh, _, err := Auth.GenRefreshTokenRecord(user.ID, ua, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create refresh token"})
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
	user.LastLogin = time.Now()
	_ = Auth.DB.Save(&user)

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
