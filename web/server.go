package web

import "github.com/gin-gonic/gin"

func NewServer() {

	r := gin.Default()

	RouteV1 := r.Group("/v1")

	RouteV1.GET("/ping", ping)

	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}
