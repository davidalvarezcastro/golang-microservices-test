package services

import (
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/models"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
)

// GetUser is a services to return a models.User
func GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	return models.GetUser(userID)
}
