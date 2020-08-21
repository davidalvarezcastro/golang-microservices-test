package app

import (
	"github.com/davidalvarezcastro/golang-microservices-test/oauth-api/src/api/controller/oauth"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
