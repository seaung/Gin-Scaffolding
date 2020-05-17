package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Demon(c *gin.Context) {
	c.JSON(http.StatusOK, "demon api")
}
