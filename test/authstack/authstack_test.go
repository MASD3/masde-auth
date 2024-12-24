package authstack_test

import (
	"testing"

	"github.com/MASD3/masde-auth/authstack"
)

func TestAuthStack(t *testing.T) {
	t.Run("register user", testRegisterUser)
	t.Run("register and verify user", testRegisterAndVerifyUser)
	t.Run("verify user with token", testVerifyUserWithToken)
	// TODO:
	// multiple users
	// try to use old token
	// duplicate users
}

func testRegisterUser(t *testing.T) {
	aS := *authstack.New()
	username := "kazuya123"
	user, err := aS.RegisterUser(username, "St0ngP4ssw0rd")
	if err != nil {
		t.Fatal("Error registering user")
	}
	if user.Username != username {
		t.Fatalf("usernames don't match: expected %s, got %s\n", username, user.Username)
	}
}

func testRegisterAndVerifyUser(t *testing.T) {
	aS := *authstack.New()
	username := "kazuya123"
	password := "St0ngP4ssw0rd"
	user, err := aS.RegisterUser(username, password)
	if err != nil {
		t.Fatal("Error registering user")
	}
	_, _, err = aS.AuthenticateWithPassword(user.Username, password)
	if err != nil {
		t.Fatalf("Error verifying user: %s", err.Error())
	}
}

func testVerifyUserWithToken(t *testing.T) {
	aS := *authstack.New()
	username := "kazuya123"
	password := "St0ngP4ssw0rd"
	user, err := aS.RegisterUser(username, password)
	if err != nil {
		t.Fatal("Error registering user")
	}
	token, _, err := aS.AuthenticateWithPassword(user.Username, password)
	if err != nil {
		t.Fatalf("Error verifying user: %s", err.Error())
	}
	_, err = aS.AuthenticateWithToken(token)
	if err != nil {
		t.Fatal("Error verifying token")
	}
}
