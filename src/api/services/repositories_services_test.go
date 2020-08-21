package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/clients/restclient"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/repositories"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo("client_id", request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "this is a description",
	}

	result, err := RepositoryService.CreateRepo("client_id", request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing"}`)),
		},
	})

	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "this is a description",
	}

	result, err := RepositoryService.CreateRepo("client_id", request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "", result.Owner)
}

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	output := make(chan repositories.CreateRespositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "this is a description",
	}
	output := make(chan repositories.CreateRespositoriesResult)

	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name:        "testing",
		Description: "this is a description",
	}
	output := make(chan repositories.CreateRespositoriesResult)

	restclient.StartMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing"}`)),
		},
	})

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.ID)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	var wg sync.WaitGroup
	input := make(chan repositories.CreateRespositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	service := reposService{}

	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)

	go func() {
		input <- repositories.CreateRespositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)

	result := <-output

	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 1, len(result.Results))
	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
}

func TestCreateReposInvalidRequest(t *testing.T) {
	request := []repositories.CreateRepoRequest{
		{},
		{Name: "  "},
	}

	result, err := RepositoryService.CreateRepos(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)
	assert.EqualValues(t, result.Results[0].Error.Status(), result.StatusCode)

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.Message())
}

func TestCreateReposOneValidOneFails(t *testing.T) {
	request := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "testing"}`)),
		},
	})

	result, err := RepositoryService.CreateRepos(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	// executes in a concurrent way => we do not know the order
	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			assert.EqualValues(t, "invalid repository name", result.Error.Message())
			continue
		}

		assert.EqualValues(t, 123, result.Response.ID)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "", result.Response.Owner)
	}
}

func TestCreateReposTwoSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		BodyText:   `{"id": 123, "name": "testing"}`,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
		},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.NotNil(t, result.Results[0].Response)
	assert.Nil(t, result.Results[0].Error)
	assert.EqualValues(t, 123, result.Results[0].Response.ID)
	assert.EqualValues(t, "testing", result.Results[0].Response.Name)
	assert.EqualValues(t, "", result.Results[0].Response.Owner)

	assert.NotNil(t, result.Results[1].Response)
	assert.Nil(t, result.Results[1].Error)
	assert.EqualValues(t, 123, result.Results[1].Response.ID)
	assert.EqualValues(t, "testing", result.Results[1].Response.Name)
	assert.EqualValues(t, "", result.Results[1].Response.Owner)
}
