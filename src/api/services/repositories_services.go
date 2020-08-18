package services

import (
	"net/http"
	"sync"

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
	CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
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
	if err := input.Validate(); err != nil {
		return nil, err
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

// CreateRepo creates many repos
func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {
	input := make(chan repositories.CreateRespositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait() // waits for the gorutines to be executed
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

// handleRepoResults handles output from channel
func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRespositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRespositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}

	output <- results
}

// createRepoConcurrent runs CreateRepo in a concurrence mode
func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRespositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRespositoriesResult{
			Error: err,
		}
		return
	}

	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRespositoriesResult{
			Error: err,
		}
		return
	}

	output <- repositories.CreateRespositoriesResult{
		Response: result,
	}
}
