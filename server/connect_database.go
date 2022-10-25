package server

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	cLog "github.com/shunta0213/test_go_auth/log"
)

func ConnectDatabase() (*sql.DB, error) {
	uri := loadConfig()
	if uri == "" {
		return nil, fmt.Errorf("Could not read .env file")
	}
	cLog.Println(uri)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func loadConfig() string {
	err := godotenv.Load(".env")

	if err != nil {
		cLog.Fatalf("Could not load .env file:", err)
		return ""
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
