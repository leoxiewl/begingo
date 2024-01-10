package server

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/list", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "list",
			})
		})

		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

	}
	return r
}
