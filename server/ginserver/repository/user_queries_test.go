package repository_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shunta0213/test_go_auth/server/ginserver/repository"
	"github.com/stretchr/testify/assert"
)

func ConnectTestDatabase() (*sql.DB, error) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", uri)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGetUserFromUserID(t *testing.T) {
	db, err := ConnectTestDatabase()
	if err != nil {
		t.Fatal("Could not connect to database", err)
	}

	tests := []struct {
		id       int64
		username string
		email    string
	}{
		{id: 1, username: "user1", email: "email1"},
		{id: 2, username: "user2", email: "email2"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("user id:", tt.id), func(t *testing.T) {
			assert := assert.New(t)

			user, err := repository.GetUserFromUserID(db, tt.id)
			if err != nil {
				t.Fatal("Could not get user", err)
			}

			assert.Equal(user.Username, tt.username)
			assert.Equal(user.Email, tt.email)
		})
	}
}

func TestGetUserFromUserName(t *testing.T) {
	db, err := ConnectTestDatabase()
	if err != nil {
		t.Fatal("Could not connect to database", err)
	}

	tests := []struct {
		id       int64
		username string
		email    string
	}{
		{id: 1, username: "user1", email: "email1"},
		{id: 2, username: "user2", email: "email2"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("user id:", tt.id), func(t *testing.T) {
			assert := assert.New(t)

			user, err := repository.GetUserFromUserName(db, tt.username)
			if err != nil {
				t.Fatal("Could not get user", err)
			}

			assert.Equal(user.Username, tt.username)
			assert.Equal(user.Email, tt.email)
		})
	}
}
func TestGetUserHashedPasswordFromUserName(t *testing.T) {
	db, err := ConnectTestDatabase()
	if err != nil {
		t.Fatal("Could not connect to database", err)
	}

	tests := []struct {
		id       int64
		password string
		username string
	}{
		{id: 1, password: "$2a$10$EJt5gEzvRnNa46ZjAoE1RuQKumCbJv.ipiNsYC4g14sn4h8CzeN/e", username: "user1"},
		{id: 2, password: "$2a$10$G7bXjXB.swnw5XpOWh994eYGRe9pxVjdgakzXwqVONg54P86fiE.K", username: "user2"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("user_id:", tt.id), func(t *testing.T) {
			assert := assert.New(t)
			user, err := repository.GetUserHashedPasswordFromUserName(db, tt.username)
			if err != nil {
				t.Fatal("Could not get user", err)
			}

			assert.Equal(tt.password, user.Password)
		})
	}
}

func TestGetUserHashedPasswordFromEmail(t *testing.T) {
	db, err := ConnectTestDatabase()
	if err != nil {
		t.Fatal("Could not connect to database", err)
	}

	tests := []struct {
		id       int64
		password string
		email    string
	}{
		{id: 1, password: "$2a$10$EJt5gEzvRnNa46ZjAoE1RuQKumCbJv.ipiNsYC4g14sn4h8CzeN/e", email: "email1"},
		{id: 2, password: "$2a$10$G7bXjXB.swnw5XpOWh994eYGRe9pxVjdgakzXwqVONg54P86fiE.K", email: "email2"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint("user_id:", tt.id), func(t *testing.T) {
			assert := assert.New(t)
			user, err := repository.GetUserHashedPasswordFromEmail(db, tt.email)
			if err != nil {
				t.Fatal("Could not get user", err)
			}

			assert.Equal(tt.password, user.Password)

			// err = password.ComparePasswordAndHash(tt.password, user.Password)
			// if err != nil {
			// 	t.Fatal("Password Not Matched")
			// }
			// t.Logf("Password Matched")
		})
	}
}

func TestCreateUser(t *testing.T) {
	db, err := ConnectTestDatabase()
	if err != nil {
		t.Fatal("Could not connect to database", err)
	}

	tests := []repository.User{
		{Username: "test1", Email: "test", Password: "test"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("user_id:", tt.ID), func(t *testing.T) {
			assert := assert.New(t)

			if err != nil {
				t.Fatal("Cannot create mock table", err)
			}

			user, err := repository.CreateUser(db, tt)
			if err != nil {
				t.Error("Could not create user", err)
			}

			assert.Equal(tt.Username, user.Username)
			assert.Equal(tt.Password, user.Password)
			assert.Equal(tt.Email, user.Email)

		})
	}
}
