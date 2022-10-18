package password

import (
	"golang.org/x/crypto/bcrypt"
)

/*
Hash function using bcrypt
*/
func HashPassword(pw string) string {
	pwByte := []byte(pw)
	hashed, err := bcrypt.GenerateFromPassword(pwByte, 10)

	if err != nil {

	}
	return string(hashed)
}

/*
Compare  a raw password and hash got from Database.

If they are matched, it return nil.
Else, it return bcrypt.ErrMismatchedHashAndPassword
*/

func ComparePasswordAndHash(pw string, hash string) error {
	pwByte := []byte(pw)
	hashByte := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashByte, pwByte)

	return err
}
