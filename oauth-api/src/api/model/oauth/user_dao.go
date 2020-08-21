package oauth

import (
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"davalv": {
			ID:       123,
			Username: "davalv",
		},
	}
)

// GetUserByUsernameAndPassword return the user id that  username and password
func GetUserByUsernameAndPassword(username string, password string) (*User, errors.APIError) {
	user := users[username]
	if user == nil {
		// return nil, errors.NewNotFoundAPIError(fmt.Sprintf("no user with username '%s'", username))
		return nil, errors.NewNotFoundAPIError("no user found with given parameters")
	}

	return user, nil
}
