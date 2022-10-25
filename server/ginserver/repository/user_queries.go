package repository

import (
	"database/sql"
)

func GetUserFromUserID(db *sql.DB, id int64) (*User, error) {
	u := User{}

	query := "SELECT * FROM users WHERE id = $1"
	err := db.QueryRow(query, id).Scan(&u)
	if err != nil {
		return nil, err
	}

	return &u, err
}

func GetUserHashedPasswordFromUserName(db *sql.DB, username string) (*User, error) {
	u := User{}

	// Query
	query := "SELECT password FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&u)
	if err != nil {
		// If no rows matches the query, this return ErrNoRows
		return nil, err
	}

	return &u, nil
}

func CreateUser(db *sql.DB, user User) (*User, error) {

	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3)"
	result, err := db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	u, err := GetUserFromUserID(db, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
