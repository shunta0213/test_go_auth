package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shunta0213/test_go_auth/server/ginserver/controllers"
	"github.com/shunta0213/test_go_auth/server/ginserver/services"
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

func TestUserController_SignUp(t *testing.T) {
	// Connect DB
	db, err := ConnectTestDatabase()
	if err != nil {
		t.Fatal("Could not connect to database", err)
	}

	// User Controller
	userController := controllers.NewUserController(services.NewUserService(db))

	// Create Gin Engine
	r := gin.Default()
	r.POST("/signup", func(ctx *gin.Context) { userController.SignUp(ctx) })

	// Mock Data
	tests := []services.SignUpDto{
		{Username: "test3", Email: "test1", Password: "pass1"},
		{Username: "test4", Email: "test2", Password: "pass2"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("Test", tt.Username), func(t *testing.T) {
			// Struct to Json and Create New Reader
			body, err := json.Marshal(tt)
			if err != nil {
				t.Fatal("Could not marshal json", err)
			}
			bodyBuffer := bytes.NewBuffer(body)

			// Send Request and Get response
			req, err := http.NewRequest("POST", "/signup", bodyBuffer)
			if err != nil {
				t.Fatal("Could not send new api request", err)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Assertion
			assert := assert.New(t)
			assert.Equal(http.StatusOK, w.Code)
		})
	}
}
