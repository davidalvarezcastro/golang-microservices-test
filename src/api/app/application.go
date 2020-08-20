package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp starts the repository api
func StartApp() {
	// log.Info("about to map the urls", "step:1", "status:pending")
	mapUrls()
	// log.Info("urls succesfully mapped", "step:1", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
