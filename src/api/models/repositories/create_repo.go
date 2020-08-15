package repositories

import (
	"strings"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

// CreateRepoRequest is the request to create a new repo
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate validates the request to create a new repo
func (r *CreateRepoRequest) Validate() errors.APIError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

// CreateRepoResponse is the response
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

// CreateReposResponse is the response from the api after creating some repos
type CreateReposResponse struct {
	StatusCode int                         `json:"status"`
	Results    []CreateRespositoriesResult `json:"results"`
}

// CreateRespositoriesResult stores a CreateRepoResponse and the error (from the api)
type CreateRespositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.APIError     `json:"error"`
}
