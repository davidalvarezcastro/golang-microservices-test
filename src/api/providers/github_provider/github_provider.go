package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/clients/restclient"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/github"
	"io/ioutil"
	"log"
	"net/http"
)

// e31e859e9407c32095d0b41aea59950e36ef47a5
const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)

}

// CreateRepo creates a repo at github domain
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse

		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json github error response body"}
		}

		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}

	return &result, nil
}
