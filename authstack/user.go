package authstack

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	uname string
	phash []byte
}

/* Generates a password hash and creates a new user struct */
func (aS *AuthStack) NewUser(uname, password string) (*User, error) {
	if _, ok := aS.users[uname]; ok {
		return nil, fmt.Errorf("Username %s is taken", uname)
	}

	phash, err := bcrypt.GenerateFromPassword([]byte(password), PWORD_HASH_COST)
	if err != nil {
		return nil, err
	}
	return &User{
		uname: uname,
		phash: phash,
	}, nil
}
