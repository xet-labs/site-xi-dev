// cntr/blog.api.go
package auth

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthApiCtrl struct {
	db  *gorm.DB
	rdb *redis.Client
}

var AuthApi = &AuthApiCtrl{}