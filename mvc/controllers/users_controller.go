package controllers

import (
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/services"
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUser is the handler to users api
func GetUser(c *gin.Context) {
	// parsing param from URL
	// c.Query("caller") if url like /users/:user_id?caller=123
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		// return the bad request to the client
		apiErr := &utils.ApplicationError{
			Message:    "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		// c.JSON(apiErr.StatusCode, apiErr)
		utils.RespondError(c, apiErr)
		return
	}

	// calling user service
	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {
		// handle error and return to the client
		// c.JSON(apiErr.StatusCode, apiErr)
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	// c.JSON(http.StatusOK, user)
	utils.Respond(c, http.StatusOK, user)
}
