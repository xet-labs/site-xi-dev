package ctrl

import "github.com/gin-gonic/gin"

// 'CustomRoutes' allows adding simple/ad-hoc routes without creating a dedicated controller
type custom struct{}

var Custom = &custom{}

func (u *custom) RouterPre(r *gin.Engine)  {}
func (u *custom) RouterCore(r *gin.Engine) {}
func (u *custom) RouterPost(r *gin.Engine) {}
