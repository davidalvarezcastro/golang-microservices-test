package services

import (
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/models"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
)

type usersService struct {
}

// UsersService variable to call users services methods
var UsersService usersService

// GetUser is a services to return a models.User
func (u *usersService) GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	// return models.UserDao.GetUser(userID)
	user, err := models.UserDao.GetUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
