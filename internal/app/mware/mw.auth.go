package mware

import (
	"fmt"
	"net/http"
	"strconv"

	model_store "xi/internal/app/model/store"
	"xi/pkg/lib/cfg"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization"})
			return
		}
		var tok string
		fmt.Sscanf(auth, "Bearer %s", &tok)
		if tok == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			return
		}

		// parse jwt
		parsed, err := jwt.Parse(tok, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("bad alg")
			}
			return []byte(cfg.Api.JwtSecret), nil
		})
		if err != nil || !parsed.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims := parsed.Claims.(jwt.RegisteredClaims)
		uid, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid subject"})
			return
		}

		// load user from DB (note: this adds a DB hit per protected route; you can cache if needed)
		db, ok := c.MustGet("DB").(*gorm.DB)
		if !ok {
			// fallback: try to get from gin engine context - for simplicity caller sets user earlier
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "server misconfig"})
			return
		}

		var user model_store.User
		if err := db.First(&user, uid).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}

		// attach user to context for handlers
		c.Set("user", &user)
		c.Next()
	}
}
