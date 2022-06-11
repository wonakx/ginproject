package main

import (
	"ginproject/biz"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		//GET
		v1.GET("/ping", biz.Pong)
		v1.GET("/query", biz.QueryString)
		v1.GET("/:path/:param", biz.PathParam)
		v1.GET("/albums", biz.GetAlbums)

	}
	v2 := router.Group("/v2")
	{
		//POST
		v2.POST("/ping", biz.PongPost)
	}

	router.Run(":8081")
}
