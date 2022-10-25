package services

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shunta0213/test_go_auth/password"
	"github.com/shunta0213/test_go_auth/server/ginserver/repository"
)

type UserService interface {
	SignUp(c *gin.Context) (*repository.User, error)
	SignIn(c *gin.Context) (*repository.User, error)
}

type userService struct {
	DB *sql.DB
}

func NewUserService(DB *sql.DB) UserService {
	return &userService{DB: DB}
}

func (s userService) SignUp(c *gin.Context) (*repository.User, error) {
	// Bind
	u := repository.User{}
	err := c.ShouldBindJSON(u)
	if err != nil {
		return nil, err
	}

	// Hash Password
	hashPass := password.HashPassword(u.Password)
	u.Password = hashPass

	// Insert
	gotUser, err := repository.CreateUser(s.DB, u)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}

func (s userService) SignIn(c *gin.Context) (*repository.User, error) {
	// Bind
	u := SingInDto{}
	if err := c.ShouldBindJSON(u); err != nil {
		return nil, err
	}

	// Get password from database
	user, err := repository.GetUserHashedPasswordFromUserName(s.DB, u.Username)
	if err != nil {
		return nil, err
	}

	// Compare raw password and hashed password from database.
	hashPass := user.Password
	err = password.ComparePasswordAndHash(u.Password, hashPass)
	if err != nil {
		return nil, err
	}

	return user, nil
}
