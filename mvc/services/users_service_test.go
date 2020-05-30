package services

import (
	"fmt"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/models"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// mocking user dao
var (
	userDaoMock usersDaoMock

	// function to overwrite GetUser method from user dao
	getUserFunction func(userID int64) (*models.User, *utils.ApplicationError)
)

// init function
func init() {
	models.UserDao = &usersDaoMock{}
}

// mocking user dao
type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*models.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v does not exist", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userID int64) (*models.User, *utils.ApplicationError) {
		return &models.User{ID: 123, FirstName: "David", LastName: "Ups", Email: "myemail@gmail.com"}, nil
	}

	user, err := UsersService.GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "David", user.FirstName)
	assert.EqualValues(t, "Ups", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}
