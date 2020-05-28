package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := GetUser(0)

	// Validation:
	// if user != nil {
	// 	t.Error("we were not expecting a user with id 0")
	// }

	// if err == nil {
	// 	t.Error("we were expecting an error when user id is 0")
	// }

	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when user is not found")
	// }

	// using testify
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 when user is not found")
	assert.EqualValues(t, "user 0 does not exist", err.Message, "we were expecting an empty message when user is not found")
	assert.EqualValues(t, "not_found", err.Code, "we were expecting a not_found code when user is not found")
}

func TestGetUser(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "David", user.FirstName)
	assert.EqualValues(t, "Ups", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)

}
