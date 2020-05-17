package v1

import (
	"github.com/seaung/Go-Scaffolding/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouters(router *gin.RouterGroup) {
	router.GET("/demon", v1.Demon)
}
