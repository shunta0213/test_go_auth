package server

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	cLog "github.com/shunta0213/test_go_auth/log"
)

func ConnectDatabase() (*sql.DB, error) {
	uri := loadConfig()
	cLog.Println(uri)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func loadConfig() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
