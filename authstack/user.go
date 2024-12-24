package authstack

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string // username (max 72 bytes)
	phash    []byte // password hash (?? bytes)
	token    []byte // session token (16 bytes)
}

/* Generates a password hash and creates a new user struct */
func (aS *AuthStack) RegisterUser(uname, password string) (*User, error) {
	if _, ok := aS.users[uname]; ok {
		return nil, fmt.Errorf("Username %s is taken", uname)
	}

	phash, err := bcrypt.GenerateFromPassword([]byte(password), PWORD_HASH_COST)
	if err != nil {
		return nil, err
	}
	user := &User{
		Username: uname,
		phash:    phash,
		token:    nil, // initial token is nil
	}
	// register the user
	aS.users[uname] = user
	return user, nil
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.phash, []byte(password))
}
