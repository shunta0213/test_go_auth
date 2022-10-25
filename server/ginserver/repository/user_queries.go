package repository

import (
	"database/sql"
	"time"
)

func GetUserFromUserID(db *sql.DB, id int64) (*User, error) {
	u := User{}

	query := "SELECT * FROM users.users WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt, &u.LastLogin)
	if err != nil {
		return nil, err
	}

	return &u, err
}

func GetUserFromUserName(db *sql.DB, username string) (*User, error) {
	u := User{}

	query := "SELECT * FROM users.users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt, &u.LastLogin)
	if err != nil {
		return nil, err
	}

	return &u, err
}

func GetUserHashedPasswordFromUserName(db *sql.DB, username string) (*User, error) {
	u := User{}

	// Query
	query := "SELECT password FROM users.users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&u.Password)
	if err != nil {
		// If no rows matches the query, this return ErrNoRows
		return nil, err
	}

	return &u, nil
}

func GetUserHashedPasswordFromEmail(db *sql.DB, email string) (*User, error) {
	u := User{}

	// Query
	query := "SELECT password FROM users.users WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&u.Password)
	if err != nil {
		// If no rows matches the query, this return ErrNoRows
		return nil, err
	}

	return &u, nil
}

func CreateUser(db *sql.DB, user *User) (*User, error) {

	query := "INSERT INTO users.users (username, email, password) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Second * 2)

	u, err := GetUserFromUserName(db, user.Username)
	if err != nil {
		return nil, err
	}

	return u, nil
}
