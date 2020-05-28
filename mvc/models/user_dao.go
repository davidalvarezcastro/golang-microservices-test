package models

import (
	"fmt"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
	"net/http"
)

// mock database
var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "David", LastName: "Ups", Email: "myemail@gmail.com"},
	}
)

// GetUser dao function to return an User
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	// user := users[userID]

	// if user == nil {
	// 	return nil, errors.New(fmt.Sprintf("user %v was not found", userID))
	// }

	// return user, nil
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
