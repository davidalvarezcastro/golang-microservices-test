package services

import (
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/models"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
	"net/http"
)

type itemsService struct {
}

// ItemsService variable to call users services methods
var ItemsService itemsService

// GetItem is a services to return a models.Item
func (i *itemsService) GetItem(itemID string) (*models.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "implement_me",
	}
}
