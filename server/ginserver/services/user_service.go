package services

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/shunta0213/test_go_auth/models"
	"github.com/shunta0213/test_go_auth/password"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserService interface {
	SignUp(c *gin.Context) (*models.User, error)
	SignIn(c *gin.Context) (*models.User, error)
}

type userService struct {
	DB *sql.DB
}

func NewUserService(DB *sql.DB) UserService {
	return &userService{DB: DB}
}

func (s userService) SignUp(c *gin.Context) (*models.User, error) {
	// Bind
	u := models.User{}
	if err := c.ShouldBindJSON(u); err != nil {
		return nil, err
	}

	// Hash Password
	hashPass := password.HashPassword(u.Password)
	u.Password = hashPass

	// Insert
	err := u.Insert(context.Background(), s.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s userService) SignIn(c *gin.Context) (*models.User, error) {
	// Bind
	u := SingInDto{}
	if err := c.ShouldBindJSON(u); err != nil {
		return nil, err
	}

	// Get password from database
	user, err := models.Users(
		qm.Select("password"),
		qm.Where("username=?", u.Username),
	).One(context.Background(), s.DB)
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
