package app

import (
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
