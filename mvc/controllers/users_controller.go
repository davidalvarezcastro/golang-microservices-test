package controllers

import (
	"encoding/json"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/services"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

// GetUser is the handler to users api
func GetUser(resp http.ResponseWriter, req *http.Request) {
	// parsing param from URL
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		// return the bad request to the client
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		apiErrJSON, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(apiErrJSON)
		return
	}
	log.Printf("about to process user id %v (%d)", userIDParam, userID)

	// calling user service
	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {
		apiErrJSON, _ := json.Marshal(apiErr)

		log.Printf("json %v", apiErr)

		// handle error and return to the client
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(apiErrJSON)
		return
	}

	// return user to client
	userJSON, _ := json.Marshal(user)
	resp.Write(userJSON)
}
