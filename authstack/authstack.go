package authstack

type AuthStack struct {
	users map[string]*User
}

func New() *AuthStack {
	return &AuthStack{
		users: make(map[string]*User),
	}
}
