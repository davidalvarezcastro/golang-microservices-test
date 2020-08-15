package repositories

import (
	"fmt"
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

	result, err := services.RepositoryService.CreateRepo(request)
	fmt.Println("asdasdsa")
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}