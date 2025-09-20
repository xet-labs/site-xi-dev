package auth

import (
	"net/http"
	"strings"
	model_store "xi/internal/app/model/store"
	"xi/pkg/app"
	"xi/pkg/lib/util"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
)


func (a *AuthApi) Signup(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Name     string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pwHash, err := util.Crypt.HashPass(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	user := model_store.User{
		Username:     strings.ToLower(req.Username),
		Email:        strings.ToLower(req.Email),
		PasswordHash: string(pwHash),
		Name:         req.Name,
	}

	db := store.Db.Cli()
	if db == nil {
		app.Err.Handle(c, app.Err.DbUnavailable, true)
		return
	}
	if err := db.Create(&user).Error; err != nil {
		app.Err.Handle(c, err, true)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}