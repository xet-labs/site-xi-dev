package auth

import (
	"errors"
	"net/http"
	"strings"

	model_store "xi/internal/app/model/store"
	"xi/pkg/app"
	"xi/pkg/lib/util"
	"xi/pkg/service/store"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a *AuthApi) Signup(c *gin.Context) {
	var req struct {
		Username        string `json:"username" binding:"min=2,max=254"`
		Email           string `json:"email" binding:"required,email,max=254"`
		Password        string `json:"password" binding:"required,min=6,max=254"`
		ConfirmPassword string `json:"confirm_password" binding:"required,min=6,max=254"`
		// Name            string `json:"name" binding:"max=254"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request, " + err.Error()})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords didnt match"})
		return
	}

	pwHash, err := util.Crypt.HashPass(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't hash password"})
		return
	}

	user := model_store.User{
		Username:     strings.ToLower(req.Username),
		Email:        strings.ToLower(req.Email),
		PasswordHash: string(pwHash),
		// Name:         req.Name,
		Config:       nil, // or util.StringPtr("{}") if you want default config
	}

	if err := store.Db.Cli().Create(&user).Error; err != nil {
		conflict := func(msg string) {
			c.JSON(http.StatusConflict, gin.H{"error": msg})
		}

		switch {
		case errors.Is(err, gorm.ErrRegistered):
			conflict("resource")
		case strings.Contains(err.Error(), "Duplicate entry"), strings.Contains(err.Error(), "unique constraint"):
			switch {
			case strings.Contains(err.Error(), "username"):
				conflict("username")
				return
			case strings.Contains(err.Error(), "email"):
				conflict("email")
				return
			}
		}
		app.Err.Handle(c, err, true)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"name":     user.Name,
		},
	})
}
