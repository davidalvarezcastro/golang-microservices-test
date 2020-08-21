package oauth

import (
	"fmt"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

// Save stores the access token (fake method, should call a database)
func (at *AccessToken) Save() errors.APIError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserID) // this shoud be a random generated string
	tokens[at.AccessToken] = at

	return nil
}

// GetAccessTokenByToken returns the info from the given token
func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.APIError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundAPIError("no access token found with given parameters")
	}

	return token, nil
}
