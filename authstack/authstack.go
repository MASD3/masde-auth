package authstack

import (
	"crypto/rand"
	"errors"
)

type AuthStack struct {
	users  map[string]*User // ASSUME PERSISTENT STORAGE
	tokens map[string]*User // ASSUME PERSISTENT STORAGE
}

func New() *AuthStack {
	return &AuthStack{
		users:  make(map[string]*User),
		tokens: make(map[string]*User),
	}
}

/* Authenticates with token returning user obj */
func (aS *AuthStack) AuthenticateWithToken(token []byte) (*User, error) {
	user, ok := aS.tokens[string(token)]
	if !ok {
		return nil, errors.New("invalid token")
	}
	return user, nil
}

/* Authenticates with uname / password pair, returning user obj and session token */
func (aS *AuthStack) AuthenticateWithPassword(uname, password string) ([]byte, *User, error) {
	user, ok := aS.users[uname]
	if !ok {
		return nil, nil, errors.New("user does not exist")
	}
	err := user.VerifyPassword(password)
	if err != nil {
		return nil, nil, err
	}
	// generate session token with crypto/rand
	token := make([]byte, SEESSION_TOKEN_LEN)
	_, err = rand.Read(token)
	if err != nil {
		return nil, nil, err
	}
	user.token = token
	// unregister/register the old/new tokens
	if user.token != nil {
		delete(aS.tokens, string(user.token))
	}
	aS.tokens[string(token)] = user

	return token, user, nil
}
