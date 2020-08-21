package oauth

import (
	"strings"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

// AccessTokenRequest stores data about the request from getting the access token
type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate validates the request
func (r *AccessTokenRequest) Validate() errors.APIError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
