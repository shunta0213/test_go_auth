package repository

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type User struct {
	ID        int64     `json:"id" toml:"id" yaml:"id"`
	Username  string    `json:"username" toml:"username" yaml:"username"`
	Email     string    `json:"email" toml:"email" yaml:"email"`
	Password  string    `json:"password" toml:"password" yaml:"password"`
	CreatedAt time.Time `json:"created_at" toml:"created_at" yaml:"created_at"`
	LastLogin null.Time `json:"last_login,omitempty" toml:"last_login" yaml:"last_login,omitempty"`
}
