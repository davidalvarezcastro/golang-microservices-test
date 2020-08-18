package app

import (
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/controllers/polo"
	"github.com/davidalvarezcastro/golang-microservices-test/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
