package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// init
func init() {
	router = gin.Default()
}

// StartApp loads all entry points
func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
