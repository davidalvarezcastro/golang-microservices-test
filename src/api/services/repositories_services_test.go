package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/clients/restclient"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repo name", err.Message())
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

	result, err := RepositoryService.CreateRepo(request)
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

	result, err := RepositoryService.CreateRepo(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
