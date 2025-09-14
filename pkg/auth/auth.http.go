// cntr/auth.go
package auth

import (
	// "xi/pkg"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthHttpCtrl struct {
	db  *gorm.DB
	rdb *redis.Client
}

// Singleton controller
var AuthHttp = &AuthHttpCtrl{
	// db:        lib.Db.GetCli(),
}
