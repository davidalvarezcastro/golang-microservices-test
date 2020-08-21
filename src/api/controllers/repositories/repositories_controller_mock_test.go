package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/repositories"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/services"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
	test "github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var (
	funcCreateRepo  func(clientID string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	funcCreateRepos func(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
)

type repoServiceMock struct{}

func (s *repoServiceMock) CreateRepo(clientID string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	return funcCreateRepo(clientID, request)
}

func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
		return &repositories.CreateRepoResponse{
			ID:    321,
			Name:  "mocked service",
			Owner: "golang",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.ID)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang", result.Owner)
}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewAPIErrFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())

}
