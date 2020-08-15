package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// APIError is the general interface for api errors
type APIError interface {
	Status() int
	Message() string
	Error() string
}

// apiError is the structure with the error data
type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

// Status returns the status
func (e *apiError) Status() int {
	return e.Astatus
}

// Message returns the message
func (e *apiError) Message() string {
	return e.Amessage
}

// Error returns the error
func (e *apiError) Error() string {
	return e.AnError
}

// NewAPIError returns a new api error
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		Astatus:  statusCode,
		Amessage: message,
	}
}

// NewAPIErrFromBytes returns an api error from bytes
func NewAPIErrFromBytes(body []byte) (APIError, error) {
	var result apiError

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}

	return &result, nil
}

// NewInternalServerError returns server error
func NewInternalServerError(message string) APIError {
	return &apiError{
		Astatus:  http.StatusInternalServerError,
		Amessage: message,
	}
}

// NewNotFoundAPIError returns not found api error
func NewNotFoundAPIError(message string) APIError {
	return &apiError{
		Astatus:  http.StatusNotFound,
		Amessage: message,
	}
}

// NewBadRequestError returns bad request error
func NewBadRequestError(message string) APIError {
	return &apiError{
		Astatus:  http.StatusBadRequest,
		Amessage: message,
	}
}
