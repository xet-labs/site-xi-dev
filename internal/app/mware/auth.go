package mware

// import (
// 	"net/http"
// 	"strings"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// )

// type AuthMw struct {

// 	once sync.Once
// 	mu   sync.RWMutex
// }

// // Middleware to verify token
// func (a *AuthMw) Authz(store *store.Store) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         // 1. correlation id set earlier
//         authHeader := c.GetHeader("Authorization")
//         if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
//             token := strings.TrimPrefix(authHeader, "Bearer ")
//             claims, err := auth.VerifyJWT(token)
//             if err == nil {
//                 c.Set("uid", claims.UserID)
//                 c.Set("scopes", claims.Scopes)
//                 c.Next()
//                 return
//             }
//             // token might be opaque (DB stored) — check DB if you use opaque
//         }

//         // 2. API key support: X-API-Key header
//         apiKey := c.GetHeader("X-API-Key")
//         if apiKey != "" {
//             // store.CheckAPIKey should check hashed key and return owner, scopes etc.
//             ownerID, scopes, ok := store.CheckAPIKey(apiKey)
//             if ok {
//                 c.Set("uid", ownerID)
//                 c.Set("scopes", scopes)
//                 c.Next()
//                 return
//             }
//         }

//         c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ok": false, "err": "unauthorized"})
//     }
// }