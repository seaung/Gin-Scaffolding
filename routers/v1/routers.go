package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(routerGroups *gin.RouterGroup) {
	InitUserRouters(routerGroups)
}
