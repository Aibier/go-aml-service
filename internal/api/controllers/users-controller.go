package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aibier/go-aml-service/internal/util/crypto"
	httperror "github.com/Aibier/go-aml-service/internal/util/http-err"

	models "github.com/Aibier/go-aml-service/internal/pkg/models/users"
	"github.com/Aibier/go-aml-service/internal/pkg/persistence"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// UserInput ...
type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
}

// GetUserByID godoc
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} users.User
// @Router /api/users/{id} [get]
// @Security Authorization Token
func GetUserByID(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Param("id")
	if user, err := s.Get(id); err != nil {
		httperror.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers godoc
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param username query string false "Username"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []users.User
// @Router /api/users [get]
// @Security Authorization Token
func GetUsers(c *gin.Context) {
	s := persistence.GetUserRepository()
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.Query(&q); err != nil {
		httperror.NewError(c, http.StatusNotFound, errors.New("users not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// CreateUser ...
func CreateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	var userInput UserInput
	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if strings.Trim(userInput.Username, " ") == "" || strings.Trim(userInput.Password, " ") == "" || strings.Trim(userInput.Lastname, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}
	existingUser, err := s.GetByUsername(userInput.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("username %s is already taken, pls login", existingUser.Username)})
		return
	}
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		Role:      models.UserRole{RoleName: userInput.Role},
	}
	if err := s.Add(&user); err != nil {
		httperror.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

// UpdateUser ...
func UpdateUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	err := c.BindJSON(&userInput)
	if err != nil {
		httperror.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	}
	if user, err := s.Get(id); err != nil {
		httperror.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		user.Username = userInput.Username
		user.Lastname = userInput.Lastname
		user.Firstname = userInput.Firstname
		user.Hash = crypto.HashAndSalt([]byte(userInput.Password))
		user.Role = models.UserRole{RoleName: userInput.Role}
		if err := s.Update(user); err != nil {
			httperror.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// DeleteUser ...
func DeleteUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.Get(id); err != nil {
		httperror.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if err := s.Delete(user); err != nil {
			httperror.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
