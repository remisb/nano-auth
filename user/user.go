package user

import "github.com/remisb/nano-auth/auth"

var userInMemStore = make(map[string]User, 0)

func init() {
	userInMemStore["remis"] = New("remis", "remis@mail.com", "remis22")
	userInMemStore["mantas"] = New("mantas", "mantas@mail.com", "mantas22")
}

func New(name, email, pass string) User {
	hash, err := auth.EncodePassword(pass)
	if err != nil {
		hash = ""
	}

	return User{
		Name:        name,
		Email:       email,
		Password:    pass,
		PassEncoded: hash,
	}
}

// Add adds a new user to the userInMemStore map.
// It takes a User object as a parameter and assigns it to the map using the user's name as the key.
// It returns an error if there was a problem adding the user.
// ```
// newUser := user.New(name, email, pass)
// user.Add(newUser)
// ```
func Add(newUser User) error {
	userInMemStore[newUser.Name] = newUser
	return nil
}

func ByNameOrEmail(name, email string) *User {
	if len(name) > 0 {
		return ByName(name)
	}

	if len(email) > 0 {
		return ByEmail(email)
	}

	return nil
}

func ByName(name string) *User {
	user := userInMemStore[name]
	return &user
}

func ByEmail(email string) *User {
	for _, user := range userInMemStore {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
