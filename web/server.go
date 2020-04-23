package web

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/hello-go/web/handlers"
)

func NewServer() {

	r := gin.Default()

	RouteV1 := r.Group("/v1")

	RouteV1.GET("/ping", handlers.PingPong)

	r.Run()
}
