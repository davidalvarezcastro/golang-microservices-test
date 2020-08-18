package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

// Mock struct contains data for our restclient mockups
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	BodyText   string
	Err        error
}

// getMockID returns a string with method and url
func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

// StartMockups starts mockups
func StartMockups() {
	enableMocks = true
}

// FlushMockups starts mockups
func FlushMockups() {
	mocks = make(map[string]*Mock)
}

// StopMockups stops mockups
func StopMockups() {
	enableMocks = false
}

// AddMockup adds a new mockup to the mockup list
func AddMockup(mock Mock) {
	mocks[getMockID(mock.HTTPMethod, mock.URL)] = &mock
}

// GetResponse returns response to the mock
func (m *Mock) GetResponse() *http.Response {
	m.Response.Body = ioutil.NopCloser(strings.NewReader(m.BodyText))
	return m.Response
}

// Post executes a post request
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMocks {
		// returns local mock without calling any external resource
		mock := mocks[getMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}

		var response *http.Response
		if mock.BodyText == "" {
			response = mock.Response
		} else {
			response = mock.GetResponse()
		}

		return response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
