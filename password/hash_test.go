package password_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shunta0213/test_go_auth/password"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func TestComparePasswordAndHash(t *testing.T) {
	tests := []string{
		"pass1",
		"pass2",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			hashedPassword := password.HashPassword(tt)
			isMatched := password.ComparePasswordAndHash(tt, hashedPassword)

			assert := assert.New(t)
			assert.Equal(nil, isMatched)
		})
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint("fail test", tt), func(t *testing.T) {
			hashedPassword := password.HashPassword(tt)
			isMatched := password.ComparePasswordAndHash(randomString(10), hashedPassword)

			assert := assert.New(t)
			assert.Equal(bcrypt.ErrMismatchedHashAndPassword, isMatched)
		})
	}
}
