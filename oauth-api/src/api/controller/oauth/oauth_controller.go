package oauth

import (
	"net/http"

	"github.com/davidalvarezcastro/golang-microservices-test/oauth-api/src/api/model/oauth"
	"github.com/davidalvarezcastro/golang-microservices-test/oauth-api/src/api/services"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateAccessToken creates an access token for our app
func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, token)

}

// GetAccessToken returns information about a given token (if exists)
func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, token)
}
