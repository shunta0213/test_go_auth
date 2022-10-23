package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	cLog "github.com/shunta0213/test_go_auth/log"
	"github.com/shunta0213/test_go_auth/server/ginserver/services"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{service}
}

func (controller userController) SignUp(c *gin.Context) {
	u, err := controller.service.SignUp(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func (controller userController) SignIn(c *gin.Context) {
	u, err := controller.service.SignIn(c)

	// Error Handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cLog.Fatal("User Not Found")
			c.AbortWithStatus(404)
			return
		}
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			cLog.Fatal("Password Mismatched")
			c.AbortWithStatus(401)
			return
		}

		// TODO When ShouldBindJson failed, what error returns?
		// When Bind Failed

		// Unexpected Error
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}
