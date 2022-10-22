package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shunta0213/test_go_auth/server/ginserver/services"
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

func (controller userController) SignUp(c *gin.Context) {}

func (controller userController) SignIn(c *gin.Context) {}
