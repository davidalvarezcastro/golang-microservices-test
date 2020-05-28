package app

import (
	"github.com/davidalvarezcastro/golang-microservices-test/mvc/controllers"
	"net/http"
)

// StartApp loads all entry points
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
