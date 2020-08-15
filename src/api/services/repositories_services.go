package services

import (
	"strings"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/config"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/providers/github_provider"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/github"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/repositories"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

var (
	// RepositoryService is a reposiroty service var
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

// CreateRepo creates a new repo in our api
func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repo name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil

}
