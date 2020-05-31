package utils

import (
	"github.com/gin-gonic/gin"
)

// Respond sends a response to the client (wrapper)
func Respond(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}

	c.JSON(status, body)
}

// RespondError sends an error to the client (wrapper)
func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}

	c.JSON(err.StatusCode, err)
}
