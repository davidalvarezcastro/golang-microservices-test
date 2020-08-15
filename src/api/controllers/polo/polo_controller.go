package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

// Marco returns a fake reponse to lete know the cloud server ours is ready
func Marco(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
