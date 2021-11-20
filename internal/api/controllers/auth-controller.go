package controllers

import (
	"errors"
	"github.com/Aibier/go-aml-service/internal/pkg/persistence"
	"github.com/Aibier/go-aml-service/pkg/crypto"
	"github.com/Aibier/go-aml-service/pkg/http-err"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginInput LoginInput
	err := c.BindJSON(&loginInput)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, errors.New("user not found"))
		log.Println(err)
	}
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
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
