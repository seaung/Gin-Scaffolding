package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/seaung/Go-Scaffolding/routers/v1"
)

var router *gin.Engine

func init() {
	router := gin.New()
	apiRouterGroup := router.Group("/api")
	v1.InitRouters(apiRouterGroup)
}

func main() {
	log.Fatal(http.ListenAndServe(":9000", router))
}
