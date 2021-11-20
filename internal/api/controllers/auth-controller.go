package controllers

import (
	"errors"
	"net/http"

	"github.com/Aibier/go-aml-service/internal/pkg/persistence"
	"github.com/Aibier/go-aml-service/pkg/crypto"
	httperror "github.com/Aibier/go-aml-service/pkg/http-err"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LoginInput ...
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login ...
func Login(c *gin.Context) {
	var loginInput LoginInput
	err := c.BindJSON(&loginInput)
	if err != nil {
		httperror.NewError(c, http.StatusBadRequest, errors.New("user not found"))
		log.Println(err)
	}
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		httperror.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			httperror.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		tokenResponse := &LoginResponse{
			Token:    token,
			Username: user.Username,
		}
		c.JSON(http.StatusOK, tokenResponse)
	}
}

//LoginResponse credential
type LoginResponse struct {
	Username string `form:"username"`
	Token    string `form:"token"`
}

//LoginCredentials credential
type LoginCredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

//LoginController contorler interface
type LoginController interface {
	Login(ctx *gin.Context) string
}
