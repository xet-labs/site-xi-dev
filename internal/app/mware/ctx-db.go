package mware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Middleware to attach DB to gin context so middleware can access it.
func DB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
