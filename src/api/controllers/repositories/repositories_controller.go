package repositories

import (
	"net/http"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/services"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"

	"github.com/davidalvarezcastro/golang-microservices-test/src/api/models/repositories"
	"github.com/gin-gonic/gin"
)

// CreateRepo creates a repo
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	cliendID := c.GetHeader("X-Client-Id")

	result, err := services.RepositoryService.CreateRepo(cliendID, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// CreateRepos creates many repos
func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)
}
