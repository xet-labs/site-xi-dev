// cntr/blog.api.go
package auth

import (
// 	"net/http"
// 	"time"

// 	model_auth "xi/internal/app/model/auth"
// 	model_db "xi/internal/app/model/db"
	lib "xi/pkg/lib"

// 	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthPkg struct {
	db  *gorm.DB
	rdb *redis.Client
}

var Auth = &AuthPkg{
	db: lib.Db.GetCli(),
}

// // Login response includes short-lived access token and profile preview.
// func (a *AuthPkg) Login(c *gin.Context) {
// 	var req model_auth.LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
// 		return
// 	}

// 	var user model_db.User
// 	if err := a.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
// 		return
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 		return
// 	}

// 	// create access token (JWT)
// 	access, err := a.GenAccessToken(&user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "token error"})
// 		return
// 	}

// 	// create refresh token (opaque), store hashed in DB
// 	rtRaw, rtHash, err := a.GenAccessToken()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation error"})
// 		return
// 	}

// 	expiresAt := time.Now().Add(h.cfg.RefreshTokenTTL)
// 	rt := model_db.User{ID: user.ID, TokenHash: rtHash, ExpiresAt: expiresAt}
// 	if err := h.db.Create(&rt).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
// 		return
// 	}

// 	// set refresh cookie (HttpOnly, Secure)
// 	h.setRefreshCookie(c, rtRaw)

// 	// response
// 	c.JSON(http.StatusOK, gin.H{
// 		"access_token": access,
// 		"expires_in":   int(h.cfg.AccessTokenTTL.Seconds()),
// 		"profile":      gin.H{"id": user.ID, "name": user.Name, "email": user.Email},
// 	})
// }
