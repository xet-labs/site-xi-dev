// cntr/auth.go
package auth

import (
	"net/http"
	"net/url"
	"xi/internal/app/lib"
	"xi/internal/app/lib/cfg"
	model_db "xi/internal/app/model/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthCtrl struct {
	db        *gorm.DB
	rdb       *redis.Client
	jwtSecret []byte
}

// Singleton controller
var Auth = &AuthCtrl{
	db:        lib.Db.GetCli(),
	jwtSecret: []byte("supersecretkey"),
}

// signup/signout
func (a *AuthCtrl) Signup(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user model_db.User

	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// Compare passwords (bcrypt)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := lib.Auth.GenToken(user.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthCtrl) ShowSignup(c *gin.Context) {
	rawUID := c.Param("uid") // @username or UID
	rawID := c.Param("id")   // auth ID or slug
	rdbKey := "/auth/" + url.QueryEscape(rawUID) + "/" + url.QueryEscape(rawID)

	// Try cache
	if lib.Web.OutCache(c, rdbKey).Html() {
		return
	}

	// Prep data
	p := cfg.Web.Pages["auths"]
	p.Rt = map[string]any{
		"url": c.Request.URL.String(),
	}

	// Cache renderer
	lib.Web.OutHtmlLyt(c, p, rdbKey)
}

func (a *AuthCtrl) ShowSignout(c *gin.Context) {}

func (a *AuthCtrl) Signout(c *gin.Context) {}

// login/logout
func (a *AuthCtrl) Logins(c *gin.Context) {
	rdbKey := c.Request.RequestURI

	// Try cache
	if lib.Web.OutCache(c, rdbKey).Html() {
		return
	}

	// Build data
	p := cfg.Web.Pages["auths"]
	p.Rt = map[string]any{
		"url": c.Request.URL.String(),
	}

	// Cache renderer
	lib.Web.OutHtmlLyt(c, p, rdbKey)
}

func (a *AuthCtrl) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user model_db.User

	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	// Compare passwords (bcrypt)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := lib.Auth.GenToken(user.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthCtrl) ShowLogin(c *gin.Context) {
	rdbKey := c.Request.RequestURI

	// Try cache
	if lib.Web.OutCache(c, rdbKey).Html() {
		return
	}

	// Prep data
	p := cfg.Web.Pages["auths"]
	p.Rt = map[string]any{
		"url": c.Request.URL.String(),
	}

	// Cache renderer
	lib.Web.OutHtmlLyt(c, p, rdbKey)
}

func (a *AuthCtrl) ShowLogout(c *gin.Context) {}

func (a *AuthCtrl) Logout(c *gin.Context) {}
